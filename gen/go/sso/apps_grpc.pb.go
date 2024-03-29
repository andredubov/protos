// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: sso/apps.proto

package ssov1

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

// AppsClient is the client API for Apps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppsClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type appsClient struct {
	cc grpc.ClientConnInterface
}

func NewAppsClient(cc grpc.ClientConnInterface) AppsClient {
	return &appsClient{cc}
}

func (c *appsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/apps.Apps/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppsServer is the server API for Apps service.
// All implementations must embed UnimplementedAppsServer
// for forward compatibility
type AppsServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedAppsServer()
}

// UnimplementedAppsServer must be embedded to have forward compatible implementations.
type UnimplementedAppsServer struct {
}

func (UnimplementedAppsServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAppsServer) mustEmbedUnimplementedAppsServer() {}

// UnsafeAppsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppsServer will
// result in compilation errors.
type UnsafeAppsServer interface {
	mustEmbedUnimplementedAppsServer()
}

func RegisterAppsServer(s grpc.ServiceRegistrar, srv AppsServer) {
	s.RegisterService(&Apps_ServiceDesc, srv)
}

func _Apps_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apps.Apps/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Apps_ServiceDesc is the grpc.ServiceDesc for Apps service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Apps_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apps.Apps",
	HandlerType: (*AppsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Apps_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/apps.proto",
}
