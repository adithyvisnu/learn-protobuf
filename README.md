# learn-protobuf
This is just an experiment for learning protobuf

## proto3 with go
https://developers.google.com/protocol-buffers/docs/overview
https://developers.google.com/protocol-buffers/docs/proto3
https://developers.google.com/protocol-buffers/docs/gotutorial

#### ref
Install protoc: https://askubuntu.com/questions/1072683/how-can-i-install-protoc-on-ubuntu-16-04
Golang documentation: https://developers.google.com/protocol-buffers/docs/reference/go-generated
Protobuf compiler: https://github.com/protocolbuffers/protobuf/releases/ 
Golang proto generator: https://godoc.org/github.com/golang/protobuf/protoc-gen-go `go get -u github.com/golang/protobuf/protoc-gen-go`
Protobuf golang packages: https://godoc.org/github.com/golang/protobuf/proto `go get github.com/golang/protobuf/proto`

##### step by step
1. write file .proto on folder proto
2. run `protoc --proto_path=proto --go_out=proto proto/user.proto` to exactly compile protobuf
3. see client.go and server.go
