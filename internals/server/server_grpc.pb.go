// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AlgorithmService_GetAllAlgorithms_FullMethodName          = "/AlgorithmService/GetAllAlgorithms"
	AlgorithmService_GetAllAlgVersionByAlgName_FullMethodName = "/AlgorithmService/GetAllAlgVersionByAlgName"
	AlgorithmService_GetAllAlgRunsByVersion_FullMethodName    = "/AlgorithmService/GetAllAlgRunsByVersion"
)

// AlgorithmServiceClient is the client API for AlgorithmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// CR Service for models
type AlgorithmServiceClient interface {
	GetAllAlgorithms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AlgorithmList, error)
	GetAllAlgVersionByAlgName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AlgorithmVersionList, error)
	GetAllAlgRunsByVersion(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AlgorithmVersionRunList, error)
}

type algorithmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlgorithmServiceClient(cc grpc.ClientConnInterface) AlgorithmServiceClient {
	return &algorithmServiceClient{cc}
}

func (c *algorithmServiceClient) GetAllAlgorithms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AlgorithmList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlgorithmList)
	err := c.cc.Invoke(ctx, AlgorithmService_GetAllAlgorithms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorithmServiceClient) GetAllAlgVersionByAlgName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AlgorithmVersionList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlgorithmVersionList)
	err := c.cc.Invoke(ctx, AlgorithmService_GetAllAlgVersionByAlgName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorithmServiceClient) GetAllAlgRunsByVersion(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AlgorithmVersionRunList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlgorithmVersionRunList)
	err := c.cc.Invoke(ctx, AlgorithmService_GetAllAlgRunsByVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgorithmServiceServer is the server API for AlgorithmService service.
// All implementations must embed UnimplementedAlgorithmServiceServer
// for forward compatibility.
//
// CR Service for models
type AlgorithmServiceServer interface {
	GetAllAlgorithms(context.Context, *emptypb.Empty) (*AlgorithmList, error)
	GetAllAlgVersionByAlgName(context.Context, *Request) (*AlgorithmVersionList, error)
	GetAllAlgRunsByVersion(context.Context, *Request) (*AlgorithmVersionRunList, error)
	mustEmbedUnimplementedAlgorithmServiceServer()
}

// UnimplementedAlgorithmServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAlgorithmServiceServer struct{}

func (UnimplementedAlgorithmServiceServer) GetAllAlgorithms(context.Context, *emptypb.Empty) (*AlgorithmList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAlgorithms not implemented")
}
func (UnimplementedAlgorithmServiceServer) GetAllAlgVersionByAlgName(context.Context, *Request) (*AlgorithmVersionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAlgVersionByAlgName not implemented")
}
func (UnimplementedAlgorithmServiceServer) GetAllAlgRunsByVersion(context.Context, *Request) (*AlgorithmVersionRunList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAlgRunsByVersion not implemented")
}
func (UnimplementedAlgorithmServiceServer) mustEmbedUnimplementedAlgorithmServiceServer() {}
func (UnimplementedAlgorithmServiceServer) testEmbeddedByValue()                          {}

// UnsafeAlgorithmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlgorithmServiceServer will
// result in compilation errors.
type UnsafeAlgorithmServiceServer interface {
	mustEmbedUnimplementedAlgorithmServiceServer()
}

func RegisterAlgorithmServiceServer(s grpc.ServiceRegistrar, srv AlgorithmServiceServer) {
	// If the following call pancis, it indicates UnimplementedAlgorithmServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AlgorithmService_ServiceDesc, srv)
}

func _AlgorithmService_GetAllAlgorithms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmServiceServer).GetAllAlgorithms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlgorithmService_GetAllAlgorithms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmServiceServer).GetAllAlgorithms(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlgorithmService_GetAllAlgVersionByAlgName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmServiceServer).GetAllAlgVersionByAlgName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlgorithmService_GetAllAlgVersionByAlgName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmServiceServer).GetAllAlgVersionByAlgName(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlgorithmService_GetAllAlgRunsByVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmServiceServer).GetAllAlgRunsByVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlgorithmService_GetAllAlgRunsByVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmServiceServer).GetAllAlgRunsByVersion(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AlgorithmService_ServiceDesc is the grpc.ServiceDesc for AlgorithmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlgorithmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AlgorithmService",
	HandlerType: (*AlgorithmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAlgorithms",
			Handler:    _AlgorithmService_GetAllAlgorithms_Handler,
		},
		{
			MethodName: "GetAllAlgVersionByAlgName",
			Handler:    _AlgorithmService_GetAllAlgVersionByAlgName_Handler,
		},
		{
			MethodName: "GetAllAlgRunsByVersion",
			Handler:    _AlgorithmService_GetAllAlgRunsByVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
