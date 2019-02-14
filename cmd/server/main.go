package main

import (
	"context"
	"log"
	"net"

	pb "github.com/renatoaquino/grpc/proto"
	"google.golang.org/grpc"
)

type helloServer struct{}

func main() {
	srv := grpc.NewServer()
	var hello helloServer
	pb.RegisterHelloServer(srv, hello)
	log.Println("starting server on port :8888")
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

func (h helloServer) Say(ctx context.Context, empty *pb.Empty) (*pb.HelloResponse, error) {
	hello := pb.HelloResponse{Message: "Hello"}
	log.Println("Say")
	return &hello, nil
}
