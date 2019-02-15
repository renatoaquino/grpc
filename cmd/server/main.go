package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/renatoaquino/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	//healthcheck on port 8080. http endpoint required by load balancers
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	go log.Fatal(http.ListenAndServe(":8080", nil))

	//begin the grpc server on port 8888
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

type helloServer struct{}

func (h helloServer) Say(ctx context.Context, empty *pb.Empty) (*pb.HelloResponse, error) {
	hello := pb.HelloResponse{Message: "Hello"}
	log.Println("Say")
	return &hello, nil
}
