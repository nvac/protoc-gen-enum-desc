package main

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/nvac/protoc-gen-enum-desc/proto/nvac"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

//go:embed template.txt
var fs embed.FS

type numberDesc struct {
	Number int32
	Desc   string
}

type nameDesc struct {
	Name string
	Desc string
}

type nameEnumType struct {
	Name     string
	EnumType *descriptorpb.EnumDescriptorProto
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	var req pluginpb.CodeGeneratorRequest
	err := proto.Unmarshal(input, &req)
	if err != nil {
		panic(err)
	}

	opts := protogen.Options{}

	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		data := getData(file)

		err = generateFile(plugin, file, data, "template.txt", "_enum_desc.pb.go")
		if err != nil {
			panic(err)
		}
	}

	out, err := proto.Marshal(plugin.Response())
	if err != nil {
		panic(err)
	}

	_, err = os.Stdout.Write(out)
	if err != nil {
		panic(err)
	}
}

func generateFile(plugin *protogen.Plugin, file *protogen.File, data map[string]any, templateFile, genFile string) error {
	t, _ := template.ParseFS(fs, templateFile)
	var buf bytes.Buffer
	_ = t.Execute(&buf, data)
	filename := file.GeneratedFilenamePrefix + genFile
	_, err := plugin.NewGeneratedFile(filename, file.GoImportPath).Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func getData(file *protogen.File) map[string]any {
	data := map[string]any{
		"Package": file.GoPackageName,
	}

	var nameEnumTypes []nameEnumType

	// top-level message
	for _, message := range file.Proto.GetMessageType() {
		getEnumTypes(&nameEnumTypes, message, message.GetName())
	}

	// top-level enum
	for _, enum := range file.Proto.GetEnumType() {
		nameEnumTypes = append(nameEnumTypes, nameEnumType{
			Name:     enum.GetName(),
			EnumType: enum,
		})
	}

	var enums []map[string]any
	for _, item := range nameEnumTypes {
		m := map[string]any{
			"EnumName": item.Name,
		}

		var numberDescs []numberDesc
		var nameDescs []nameDesc
		for _, value := range item.EnumType.Value {
			v := proto.GetExtension(value.Options, nvac.E_EnumDesc)
			numberDescs = append(numberDescs, numberDesc{
				Number: value.GetNumber(),
				Desc:   v.(string),
			})

			nameDescs = append(nameDescs, nameDesc{
				Name: value.GetName(),
				Desc: v.(string),
			})
		}
		m["NumberDescs"] = numberDescs
		m["NameDescs"] = nameDescs
		enums = append(enums, m)
	}

	data["Enums"] = enums

	return data
}

func getEnumTypes(nameEnumTypes *[]nameEnumType, descriptor *descriptorpb.DescriptorProto, name string) {
	for _, enum := range descriptor.GetEnumType() {
		*nameEnumTypes = append(*nameEnumTypes, nameEnumType{
			Name:     fmt.Sprintf("%s_%s", name, enum.GetName()),
			EnumType: enum,
		})
	}

	for _, nestedDescriptor := range descriptor.GetNestedType() {
		getEnumTypes(nameEnumTypes, nestedDescriptor, fmt.Sprintf("%s_%s", name, nestedDescriptor.GetName()))
	}
}
