// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PreludeClient is the client API for Prelude service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PreludeClient interface {
	//ping 服务check接口
	Ping(ctx context.Context, in *STPreludePingReq, opts ...grpc.CallOption) (*STPreludePingRsp, error)
}

type preludeClient struct {
	cc grpc.ClientConnInterface
}

func NewPreludeClient(cc grpc.ClientConnInterface) PreludeClient {
	return &preludeClient{cc}
}

func (c *preludeClient) Ping(ctx context.Context, in *STPreludePingReq, opts ...grpc.CallOption) (*STPreludePingRsp, error) {
	out := new(STPreludePingRsp)
	err := c.cc.Invoke(ctx, "/prelude.Prelude/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreludeServer is the server API for Prelude service.
// All implementations must embed UnimplementedPreludeServer
// for forward compatibility
type PreludeServer interface {
	//ping 服务check接口
	Ping(context.Context, *STPreludePingReq) (*STPreludePingRsp, error)
	mustEmbedUnimplementedPreludeServer()
}

// UnimplementedPreludeServer must be embedded to have forward compatible implementations.
type UnimplementedPreludeServer struct {
}

func (UnimplementedPreludeServer) Ping(context.Context, *STPreludePingReq) (*STPreludePingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPreludeServer) mustEmbedUnimplementedPreludeServer() {}

// UnsafePreludeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PreludeServer will
// result in compilation errors.
type UnsafePreludeServer interface {
	mustEmbedUnimplementedPreludeServer()
}

func RegisterPreludeServer(s grpc.ServiceRegistrar, srv PreludeServer) {
	s.RegisterService(&Prelude_ServiceDesc, srv)
}

func _Prelude_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(STPreludePingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreludeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prelude.Prelude/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreludeServer).Ping(ctx, req.(*STPreludePingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Prelude_ServiceDesc is the grpc.ServiceDesc for Prelude service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Prelude_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prelude.Prelude",
	HandlerType: (*PreludeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Prelude_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/prelude.proto",
}