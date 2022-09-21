// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: grpc.proto

package grpc

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

// GrpcClient is the client API for Grpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthReply, error)
	AddTodo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoReply, error)
}

type grpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcClient(cc grpc.ClientConnInterface) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Grpc/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/Grpc/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthReply, error) {
	out := new(HealthReply)
	err := c.cc.Invoke(ctx, "/Grpc/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) AddTodo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoReply, error) {
	out := new(TodoReply)
	err := c.cc.Invoke(ctx, "/Grpc/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServer is the server API for Grpc service.
// All implementations must embed UnimplementedGrpcServer
// for forward compatibility
type GrpcServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Sum(context.Context, *SumRequest) (*SumReply, error)
	Health(context.Context, *HealthRequest) (*HealthReply, error)
	AddTodo(context.Context, *TodoRequest) (*TodoReply, error)
	mustEmbedUnimplementedGrpcServer()
}

// UnimplementedGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcServer struct {
}

func (UnimplementedGrpcServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGrpcServer) Sum(context.Context, *SumRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedGrpcServer) Health(context.Context, *HealthRequest) (*HealthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedGrpcServer) AddTodo(context.Context, *TodoRequest) (*TodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (UnimplementedGrpcServer) mustEmbedUnimplementedGrpcServer() {}

// UnsafeGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcServer will
// result in compilation errors.
type UnsafeGrpcServer interface {
	mustEmbedUnimplementedGrpcServer()
}

func RegisterGrpcServer(s grpc.ServiceRegistrar, srv GrpcServer) {
	s.RegisterService(&Grpc_ServiceDesc, srv)
}

func _Grpc_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Grpc/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Grpc/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Grpc/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Grpc/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).AddTodo(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Grpc_ServiceDesc is the grpc.ServiceDesc for Grpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Grpc_SayHello_Handler,
		},
		{
			MethodName: "Sum",
			Handler:    _Grpc_Sum_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Grpc_Health_Handler,
		},
		{
			MethodName: "AddTodo",
			Handler:    _Grpc_AddTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
