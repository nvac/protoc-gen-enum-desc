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
var fsys embed.FS

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
	proto.Unmarshal(input, &req)

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

		fs, _ := template.ParseFS(fsys, "template.txt")

		var buf bytes.Buffer
		_ = fs.Execute(&buf, data)

		filename := file.GeneratedFilenamePrefix + ".enum_desc.pb.go"

		plugin.NewGeneratedFile(filename, file.GoImportPath).Write(buf.Bytes())
	}

	out, err := proto.Marshal(plugin.Response())
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(out)
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
