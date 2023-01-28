package main

import (
	"context"
	"grpc-demo/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":8090", grpc.WithInsecure())
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	_ = SayHello(client)
}

func SayHello(client proto.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &proto.HelloRequest{Name: "this is a test name"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
