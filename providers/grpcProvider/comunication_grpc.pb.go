// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcProvider

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

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Services_ConnectClient, error)
}

type servicesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesClient(cc grpc.ClientConnInterface) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Services_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Services_ServiceDesc.Streams[0], "/grpcProvider.Services/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicesConnectClient{stream}
	return x, nil
}

type Services_ConnectClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type servicesConnectClient struct {
	grpc.ClientStream
}

func (x *servicesConnectClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *servicesConnectClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServicesServer is the server API for Services service.
// All implementations must embed UnimplementedServicesServer
// for forward compatibility
type ServicesServer interface {
	Connect(Services_ConnectServer) error
	mustEmbedUnimplementedServicesServer()
}

// UnimplementedServicesServer must be embedded to have forward compatible implementations.
type UnimplementedServicesServer struct {
}

func (UnimplementedServicesServer) Connect(Services_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedServicesServer) mustEmbedUnimplementedServicesServer() {}

// UnsafeServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServer will
// result in compilation errors.
type UnsafeServicesServer interface {
	mustEmbedUnimplementedServicesServer()
}

func RegisterServicesServer(s grpc.ServiceRegistrar, srv ServicesServer) {
	s.RegisterService(&Services_ServiceDesc, srv)
}

func _Services_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServicesServer).Connect(&servicesConnectServer{stream})
}

type Services_ConnectServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type servicesConnectServer struct {
	grpc.ServerStream
}

func (x *servicesConnectServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *servicesConnectServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Services_ServiceDesc is the grpc.ServiceDesc for Services service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Services_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcProvider.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Services_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "comunication.proto",
}
