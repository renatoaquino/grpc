package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	pb "github.com/renatoaquino/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	addr := "0.0.0.0:8888"
	if flag.NArg() == 1 {
		addr = flag.Arg(0)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := pb.NewHelloClient(conn)

	response, err := client.Say(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not request to backend: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("server said: %s\n", response.Message)
}
