package main

import (
	"bytes"
	"embed"
	_ "embed"
	"io"
	"os"
	"text/template"

	"github.com/nvac/protoc-gen-enum-desc/proto/nvac"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
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

	for _, enum := range file.Proto.GetEnumType() {
		data["EnumName"] = enum.GetName()

		var numberDescs []numberDesc
		var nameDescs []nameDesc
		for _, value := range enum.Value {
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
		data["NumberDescs"] = numberDescs
		data["NameDescs"] = nameDescs
	}

	return data
}
