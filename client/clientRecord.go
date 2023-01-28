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
	_ = SayRecord(client, &proto.HelloRequest{Name: "CICI"})
}

func SayRecord(client proto.GreeterClient, r *proto.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()
	log.Printf("resp err:%v", resp)
	return nil
}
