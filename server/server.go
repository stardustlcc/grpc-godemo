package main

import (
	"context"
	"grpc-demo/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreeterServers struct {
	proto.UnimplementedGreeterServer
}

func (s *GreeterServers) SayHello(ctx context.Context, r *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: r.Name}, nil
}

func (s *GreeterServers) SayList(r *proto.HelloRequest, stream proto.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&proto.HelloReply{Message: "hello cici"})
	}
	return nil
}

func (s *GreeterServers) SayRecord(stream proto.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			message := &proto.HelloReply{Message: "say.record"}
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}
		log.Printf("resp:%v", resp)
	}
	return nil
}

func (s *GreeterServers) SayRoute(stream proto.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&proto.HelloReply{Message: "say.route"})
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		log.Printf("resp:%v", resp)
	}
}

func main() {
	server := grpc.NewServer()
	proto.RegisterGreeterServer(server, &GreeterServers{})
	lis, _ := net.Listen("tcp", ":8090")
	server.Serve(lis)
}
