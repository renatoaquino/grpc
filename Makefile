.DEFAULT_GOAL:= default

default: clean
	go build -o client ./cmd/client
	go build -o server ./cmd/server
clean:
	rm -f client server server-linux

deps:
	go get github.com/golang/protobuf/protoc-gen-go
	go mod tidy

proto:
	protoc --go_out=plugins=grpc:. proto/*.proto

docker:
	GOOS=linux GOARCH=amd64 go build -o server-linux ./cmd/server
	docker build -t grpc:latest .