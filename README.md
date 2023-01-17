# protoc-gen-enum-desc

Generate custom description for proto Enum.

### Install

`go install github.com/nvac/protoc-gen-enum-desc@latest`

### Usage

1. copy `./proto/nvac/enum_desc.proto` to your proto buffer import path

2. import option `nvac.enum_desc` from `nvac/enum_desc.proto`
    ````protobuf
    syntax = "proto3";
    
    package example;
    
    import "google/protobuf/descriptor.proto";
    import "nvac/enum_desc.proto";
    
    option go_package = ".;example";
    
    message User {
      Status status = 1;
    }
    
    enum Status {
      Status_Unspecified = 0 [(nvac.enum_desc) = "UNSPECIFIED"];
      Status_Active = 1 [(nvac.enum_desc) = "ACTIVE"];
      Status_Deleted = 2 [(nvac.enum_desc) = "DELETED"];
    }
    ````

3. compile with `--enum-desc_out`
    ````shell
    protoc --proto_path=. \
        --proto_path=./proto \
        --go_out=paths=source_relative:. \
        --enum-desc_out=paths=source_relative:. \
        ./example/user.proto
    ````