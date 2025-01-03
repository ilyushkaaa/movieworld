// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

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

// AuthMakerClient is the client API for AuthMaker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthMakerClient interface {
	Login(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*User, error)
	Register(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*User, error)
	CreateSession(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error)
	GetSession(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error)
	DeleteSession(ctx context.Context, in *Token, opts ...grpc.CallOption) (*IsDeleted, error)
}

type authMakerClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthMakerClient(cc grpc.ClientConnInterface) AuthMakerClient {
	return &authMakerClient{cc}
}

func (c *authMakerClient) Login(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.AuthMaker/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMakerClient) Register(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.AuthMaker/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMakerClient) CreateSession(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth.AuthMaker/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMakerClient) GetSession(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/auth.AuthMaker/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMakerClient) DeleteSession(ctx context.Context, in *Token, opts ...grpc.CallOption) (*IsDeleted, error) {
	out := new(IsDeleted)
	err := c.cc.Invoke(ctx, "/auth.AuthMaker/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthMakerServer is the server API for AuthMaker service.
// All implementations must embed UnimplementedAuthMakerServer
// for forward compatibility
type AuthMakerServer interface {
	Login(context.Context, *AuthData) (*User, error)
	Register(context.Context, *AuthData) (*User, error)
	CreateSession(context.Context, *User) (*Token, error)
	GetSession(context.Context, *Token) (*Session, error)
	DeleteSession(context.Context, *Token) (*IsDeleted, error)
	mustEmbedUnimplementedAuthMakerServer()
}

// UnimplementedAuthMakerServer must be embedded to have forward compatible implementations.
type UnimplementedAuthMakerServer struct {
}

func (UnimplementedAuthMakerServer) Login(context.Context, *AuthData) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthMakerServer) Register(context.Context, *AuthData) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthMakerServer) CreateSession(context.Context, *User) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedAuthMakerServer) GetSession(context.Context, *Token) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedAuthMakerServer) DeleteSession(context.Context, *Token) (*IsDeleted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedAuthMakerServer) mustEmbedUnimplementedAuthMakerServer() {}

// UnsafeAuthMakerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthMakerServer will
// result in compilation errors.
type UnsafeAuthMakerServer interface {
	mustEmbedUnimplementedAuthMakerServer()
}

func RegisterAuthMakerServer(s grpc.ServiceRegistrar, srv AuthMakerServer) {
	s.RegisterService(&AuthMaker_ServiceDesc, srv)
}

func _AuthMaker_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMakerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthMaker/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMakerServer).Login(ctx, req.(*AuthData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMaker_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMakerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthMaker/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMakerServer).Register(ctx, req.(*AuthData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMaker_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMakerServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthMaker/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMakerServer).CreateSession(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMaker_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMakerServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthMaker/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMakerServer).GetSession(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMaker_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMakerServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthMaker/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMakerServer).DeleteSession(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthMaker_ServiceDesc is the grpc.ServiceDesc for AuthMaker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthMaker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthMaker",
	HandlerType: (*AuthMakerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthMaker_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthMaker_Register_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _AuthMaker_CreateSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _AuthMaker_GetSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _AuthMaker_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
