.PHONY: init
init:
	protoc --proto_path=./proto --go_out=paths=source_relative:./proto ./proto/nvac/enum_desc.proto

.PHONY: example
example:
	go build
	protoc --plugin=protoc-gen-enum-desc=./protoc-gen-enum-desc \
	       --proto_path=. \
	       --proto_path=./proto \
 	       --go_out=paths=source_relative:. \
 	       --enum-desc_out=paths=source_relative:. \
	       ./example/user.proto