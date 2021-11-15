// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SphinxAgentClient is the client API for SphinxAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SphinxAgentClient interface {
	// 交易状态查询
	GetTxStatus(ctx context.Context, in *GetTxStatusRequest, opts ...grpc.CallOption) (*GetTxStatusResponse, error)
	// 余额查询
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*AccountBalance, error)
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type sphinxAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewSphinxAgentClient(cc grpc.ClientConnInterface) SphinxAgentClient {
	return &sphinxAgentClient{cc}
}

func (c *sphinxAgentClient) GetTxStatus(ctx context.Context, in *GetTxStatusRequest, opts ...grpc.CallOption) (*GetTxStatusResponse, error) {
	out := new(GetTxStatusResponse)
	err := c.cc.Invoke(ctx, "/sphinx.agent.v1.SphinxAgent/GetTxStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxAgentClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*AccountBalance, error) {
	out := new(AccountBalance)
	err := c.cc.Invoke(ctx, "/sphinx.agent.v1.SphinxAgent/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxAgentClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/sphinx.agent.v1.SphinxAgent/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SphinxAgentServer is the server API for SphinxAgent service.
// All implementations must embed UnimplementedSphinxAgentServer
// for forward compatibility
type SphinxAgentServer interface {
	// 交易状态查询
	GetTxStatus(context.Context, *GetTxStatusRequest) (*GetTxStatusResponse, error)
	// 余额查询
	GetBalance(context.Context, *GetBalanceRequest) (*AccountBalance, error)
	// Method Version
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	mustEmbedUnimplementedSphinxAgentServer()
}

// UnimplementedSphinxAgentServer must be embedded to have forward compatible implementations.
type UnimplementedSphinxAgentServer struct {
}

func (UnimplementedSphinxAgentServer) GetTxStatus(context.Context, *GetTxStatusRequest) (*GetTxStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxStatus not implemented")
}
func (UnimplementedSphinxAgentServer) GetBalance(context.Context, *GetBalanceRequest) (*AccountBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedSphinxAgentServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedSphinxAgentServer) mustEmbedUnimplementedSphinxAgentServer() {}

// UnsafeSphinxAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SphinxAgentServer will
// result in compilation errors.
type UnsafeSphinxAgentServer interface {
	mustEmbedUnimplementedSphinxAgentServer()
}

func RegisterSphinxAgentServer(s grpc.ServiceRegistrar, srv SphinxAgentServer) {
	s.RegisterService(&SphinxAgent_ServiceDesc, srv)
}

func _SphinxAgent_GetTxStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxAgentServer).GetTxStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.agent.v1.SphinxAgent/GetTxStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxAgentServer).GetTxStatus(ctx, req.(*GetTxStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxAgent_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxAgentServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.agent.v1.SphinxAgent/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxAgentServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxAgent_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxAgentServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.agent.v1.SphinxAgent/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxAgentServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SphinxAgent_ServiceDesc is the grpc.ServiceDesc for SphinxAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SphinxAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.agent.v1.SphinxAgent",
	HandlerType: (*SphinxAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTxStatus",
			Handler:    _SphinxAgent_GetTxStatus_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _SphinxAgent_GetBalance_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _SphinxAgent_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/sphinx-agent.proto",
}