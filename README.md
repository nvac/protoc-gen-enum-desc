# protoc-gen-enum-desc

Generates custom description for proto Enum.

1. copy `./proto/nvac/enum_desc.proto` to your protobuffer import path

2. import option `nvac.enum_desc` from `nvac/enum_desc.proto`
````protobuf
syntax = "proto3";

package example;

import "google/protobuf/descriptor.proto";
import "nvac/enum_desc.proto";

option go_package = ".;example";

message User {
  string username = 1;
  Status status = 2;
}

enum Status {
  _ = 0;
  Active = 1 [(nvac.enum_desc) = "ACTIVE"];
  Deleted = 2 [(nvac.enum_desc) = "DELETED"];
}
````
3. `go install github.com/nvac/protoc-gen-enum-desc@latest`

4. protoc with `--enum-desc_out`
````shell
protoc --plugin=protoc-gen-enum-desc=./protoc-gen-enum-desc \
    --proto_path=. \
    --proto_path=./proto \
    --go_out=paths=source_relative:. \
    --enum-desc_out=paths=source_relative:. \
    ./example/user.proto
````