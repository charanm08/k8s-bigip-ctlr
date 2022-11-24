// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: controller/agent.proto

package controller

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

// EndPointServiceClient is the client API for EndPointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndPointServiceClient interface {
	// Sends a greeting
	ProcessEndpoints(ctx context.Context, in *ClusterAgentRequest, opts ...grpc.CallOption) (*ServerResponse, error)
}

type endPointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEndPointServiceClient(cc grpc.ClientConnInterface) EndPointServiceClient {
	return &endPointServiceClient{cc}
}

func (c *endPointServiceClient) ProcessEndpoints(ctx context.Context, in *ClusterAgentRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/controller.EndPointService/ProcessEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndPointServiceServer is the server API for EndPointService service.
// All implementations must embed UnimplementedEndPointServiceServer
// for forward compatibility
type EndPointServiceServer interface {
	// Sends a greeting
	ProcessEndpoints(context.Context, *ClusterAgentRequest) (*ServerResponse, error)
	mustEmbedUnimplementedEndPointServiceServer()
}

// UnimplementedEndPointServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEndPointServiceServer struct {
}

func (UnimplementedEndPointServiceServer) ProcessEndpoints(context.Context, *ClusterAgentRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessEndpoints not implemented")
}
func (UnimplementedEndPointServiceServer) mustEmbedUnimplementedEndPointServiceServer() {}

// UnsafeEndPointServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndPointServiceServer will
// result in compilation errors.
type UnsafeEndPointServiceServer interface {
	mustEmbedUnimplementedEndPointServiceServer()
}

func RegisterEndPointServiceServer(s grpc.ServiceRegistrar, srv EndPointServiceServer) {
	s.RegisterService(&EndPointService_ServiceDesc, srv)
}

func _EndPointService_ProcessEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndPointServiceServer).ProcessEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.EndPointService/ProcessEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndPointServiceServer).ProcessEndpoints(ctx, req.(*ClusterAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EndPointService_ServiceDesc is the grpc.ServiceDesc for EndPointService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndPointService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controller.EndPointService",
	HandlerType: (*EndPointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessEndpoints",
			Handler:    _EndPointService_ProcessEndpoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/agent.proto",
}
