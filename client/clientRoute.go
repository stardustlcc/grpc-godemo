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
	_ = SayRoute(client, &proto.HelloRequest{Name: "ha ha ha ha"})
}

func SayRoute(client proto.GreeterClient, r *proto.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp err: %v", resp)
	}
	_ = stream.CloseSend()
	return nil
}
