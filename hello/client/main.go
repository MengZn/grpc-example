package main

import (

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-example/proto/hello"
	"log"
	"google.golang.org/grpc/grpclog"
)

const (
	//gRPC ip port
	Address = "127.0.0.1:8080"
)

func main() {
	//定義grpc連結的位置與方法
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	// 初始化客户端
	client := hello.NewHelloClient(conn)

	// 使用hello Request的結構
	req := &hello.HelloRequest{Name: "gRPC"}
	// 將上述結構丟到say hello的function 傳送出去
	res, err := client.SayHello(context.Background(), req)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Message)
}