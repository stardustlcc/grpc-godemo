package main

import (
	"context"
	"grpc-demo/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":8090", grpc.WithInsecure())
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	_ = SayList(client, &proto.HelloRequest{})
}

func SayList(client proto.GreeterClient, r *proto.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: %v", resp)
	}
	return nil
}
