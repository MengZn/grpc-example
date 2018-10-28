package main

import (
	"net"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc"
	"grpc-example/proto/hello"
	"fmt"
	"context"
	"log"
	"google.golang.org/grpc/reflection"
)

var address = "0.0.0.0:8080"

// 實現HelloServer的type
type helloService struct{}

// 實現HelloServer SayHello function
func (h helloService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		grpclog.Fatalf("Unable to listen on 8080 port: %v", err)
	}
	server := grpc.NewServer()
	hello.RegisterHelloServer(server, helloService{})
	reflection.Register(server)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
