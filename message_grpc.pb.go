// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// DockerServiceClient is the client API for DockerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerServiceClient interface {
	GetDockerVersion(ctx context.Context, in *Client, opts ...grpc.CallOption) (*DockerVersion, error)
}

type dockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerServiceClient(cc grpc.ClientConnInterface) DockerServiceClient {
	return &dockerServiceClient{cc}
}

func (c *dockerServiceClient) GetDockerVersion(ctx context.Context, in *Client, opts ...grpc.CallOption) (*DockerVersion, error) {
	out := new(DockerVersion)
	err := c.cc.Invoke(ctx, "/main.DockerService/GetDockerVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerServiceServer is the server API for DockerService service.
// All implementations must embed UnimplementedDockerServiceServer
// for forward compatibility
type DockerServiceServer interface {
	GetDockerVersion(context.Context, *Client) (*DockerVersion, error)
	mustEmbedUnimplementedDockerServiceServer()
}

// UnimplementedDockerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDockerServiceServer struct {
}

func (UnimplementedDockerServiceServer) GetDockerVersion(context.Context, *Client) (*DockerVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDockerVersion not implemented")
}
func (UnimplementedDockerServiceServer) mustEmbedUnimplementedDockerServiceServer() {}

// UnsafeDockerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerServiceServer will
// result in compilation errors.
type UnsafeDockerServiceServer interface {
	mustEmbedUnimplementedDockerServiceServer()
}

func RegisterDockerServiceServer(s grpc.ServiceRegistrar, srv DockerServiceServer) {
	s.RegisterService(&DockerService_ServiceDesc, srv)
}

func _DockerService_GetDockerVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetDockerVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.DockerService/GetDockerVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetDockerVersion(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

// DockerService_ServiceDesc is the grpc.ServiceDesc for DockerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.DockerService",
	HandlerType: (*DockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDockerVersion",
			Handler:    _DockerService_GetDockerVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
