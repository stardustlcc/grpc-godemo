// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protolist.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayList(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayListClient, error)
	SayRecord(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayRecordClient, error)
	SayRoute(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayRouteClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/proto.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayList(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/proto.Greeter/SayList", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayListClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayListClient struct {
	grpc.ClientStream
}

func (x *greeterSayListClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayRecord(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayRecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/proto.Greeter/SayRecord", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayRecordClient{stream}
	return x, nil
}

type Greeter_SayRecordClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayRecordClient struct {
	grpc.ClientStream
}

func (x *greeterSayRecordClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayRecordClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayRoute(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], "/proto.Greeter/SayRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayRouteClient{stream}
	return x, nil
}

type Greeter_SayRouteClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayRouteClient struct {
	grpc.ClientStream
}

func (x *greeterSayRouteClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayRouteClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayList(*HelloRequest, Greeter_SayListServer) error
	SayRecord(Greeter_SayRecordServer) error
	SayRoute(Greeter_SayRouteServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) SayList(*HelloRequest, Greeter_SayListServer) error {
	return status.Errorf(codes.Unimplemented, "method SayList not implemented")
}
func (UnimplementedGreeterServer) SayRecord(Greeter_SayRecordServer) error {
	return status.Errorf(codes.Unimplemented, "method SayRecord not implemented")
}
func (UnimplementedGreeterServer) SayRoute(Greeter_SayRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method SayRoute not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayList(m, &greeterSayListServer{stream})
}

type Greeter_SayListServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayListServer struct {
	grpc.ServerStream
}

func (x *greeterSayListServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_SayRecord_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayRecord(&greeterSayRecordServer{stream})
}

type Greeter_SayRecordServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayRecordServer struct {
	grpc.ServerStream
}

func (x *greeterSayRecordServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayRecordServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayRoute(&greeterSayRouteServer{stream})
}

type Greeter_SayRouteServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayRouteServer struct {
	grpc.ServerStream
}

func (x *greeterSayRouteServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayRouteServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayList",
			Handler:       _Greeter_SayList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayRecord",
			Handler:       _Greeter_SayRecord_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayRoute",
			Handler:       _Greeter_SayRoute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protolist.proto",
}
