//Package proto has the protocol buffer related files
//To generate the service go code you'll require
//protobuf - brew install protobuf
//protoc-gen-go - go get github.com/golang/protobuf/protoc-gen-go
//
//protoc --go_out=plugins=grpc:. *.proto
//
//None of that is needed to run the application, only if the .proto files are changed.
package proto
