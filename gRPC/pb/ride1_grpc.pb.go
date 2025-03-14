// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: ride1.proto

package app

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Rides_Start_FullMethodName = "/rides.Rides/Start"
)

// RidesClient is the client API for Rides service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RidesClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
}

type ridesClient struct {
	cc grpc.ClientConnInterface
}

func NewRidesClient(cc grpc.ClientConnInterface) RidesClient {
	return &ridesClient{cc}
}

func (c *ridesClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, Rides_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RidesServer is the server API for Rides service.
// All implementations must embed UnimplementedRidesServer
// for forward compatibility.
type RidesServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	mustEmbedUnimplementedRidesServer()
}

// UnimplementedRidesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRidesServer struct{}

func (UnimplementedRidesServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedRidesServer) mustEmbedUnimplementedRidesServer() {}
func (UnimplementedRidesServer) testEmbeddedByValue()               {}

// UnsafeRidesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RidesServer will
// result in compilation errors.
type UnsafeRidesServer interface {
	mustEmbedUnimplementedRidesServer()
}

func RegisterRidesServer(s grpc.ServiceRegistrar, srv RidesServer) {
	// If the following call pancis, it indicates UnimplementedRidesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rides_ServiceDesc, srv)
}

func _Rides_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RidesServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rides_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RidesServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rides_ServiceDesc is the grpc.ServiceDesc for Rides service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rides_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rides.Rides",
	HandlerType: (*RidesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Rides_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ride1.proto",
}
