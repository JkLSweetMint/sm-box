// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: src/transport/proto/src/users-service/access_system-service.proto

package __

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

// AccessSystemServiceClient is the client API for AccessSystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessSystemServiceClient interface {
	CheckUserAccess(ctx context.Context, in *AccessSystemCheckUserAccessRequest, opts ...grpc.CallOption) (*AccessSystemCheckUserAccessResponse, error)
}

type accessSystemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessSystemServiceClient(cc grpc.ClientConnInterface) AccessSystemServiceClient {
	return &accessSystemServiceClient{cc}
}

func (c *accessSystemServiceClient) CheckUserAccess(ctx context.Context, in *AccessSystemCheckUserAccessRequest, opts ...grpc.CallOption) (*AccessSystemCheckUserAccessResponse, error) {
	out := new(AccessSystemCheckUserAccessResponse)
	err := c.cc.Invoke(ctx, "/grpc_service.AccessSystemService/CheckUserAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessSystemServiceServer is the server API for AccessSystemService service.
// All implementations must embed UnimplementedAccessSystemServiceServer
// for forward compatibility
type AccessSystemServiceServer interface {
	CheckUserAccess(context.Context, *AccessSystemCheckUserAccessRequest) (*AccessSystemCheckUserAccessResponse, error)
	mustEmbedUnimplementedAccessSystemServiceServer()
}

// UnimplementedAccessSystemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccessSystemServiceServer struct {
}

func (UnimplementedAccessSystemServiceServer) CheckUserAccess(context.Context, *AccessSystemCheckUserAccessRequest) (*AccessSystemCheckUserAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAccess not implemented")
}
func (UnimplementedAccessSystemServiceServer) mustEmbedUnimplementedAccessSystemServiceServer() {}

// UnsafeAccessSystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessSystemServiceServer will
// result in compilation errors.
type UnsafeAccessSystemServiceServer interface {
	mustEmbedUnimplementedAccessSystemServiceServer()
}

func RegisterAccessSystemServiceServer(s grpc.ServiceRegistrar, srv AccessSystemServiceServer) {
	s.RegisterService(&AccessSystemService_ServiceDesc, srv)
}

func _AccessSystemService_CheckUserAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessSystemCheckUserAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessSystemServiceServer).CheckUserAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_service.AccessSystemService/CheckUserAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessSystemServiceServer).CheckUserAccess(ctx, req.(*AccessSystemCheckUserAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessSystemService_ServiceDesc is the grpc.ServiceDesc for AccessSystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessSystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_service.AccessSystemService",
	HandlerType: (*AccessSystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUserAccess",
			Handler:    _AccessSystemService_CheckUserAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/transport/proto/src/users-service/access_system-service.proto",
}
