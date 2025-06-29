// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: users.proto

package users

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
	UsersService_SaveSecurityImage_FullMethodName = "/users.UsersService/SaveSecurityImage"
	UsersService_GetSecurityImage_FullMethodName  = "/users.UsersService/GetSecurityImage"
	UsersService_CreateUser_FullMethodName        = "/users.UsersService/CreateUser"
	UsersService_GetUserByEmail_FullMethodName    = "/users.UsersService/GetUserByEmail"
)

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	SaveSecurityImage(ctx context.Context, in *SaveSecurityImageRequest, opts ...grpc.CallOption) (*SaveSecurityImageResponse, error)
	GetSecurityImage(ctx context.Context, in *GetSecurityImageRequest, opts ...grpc.CallOption) (*GetSecurityImageResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) SaveSecurityImage(ctx context.Context, in *SaveSecurityImageRequest, opts ...grpc.CallOption) (*SaveSecurityImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveSecurityImageResponse)
	err := c.cc.Invoke(ctx, UsersService_SaveSecurityImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetSecurityImage(ctx context.Context, in *GetSecurityImageRequest, opts ...grpc.CallOption) (*GetSecurityImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSecurityImageResponse)
	err := c.cc.Invoke(ctx, UsersService_GetSecurityImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UsersService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserByEmailResponse)
	err := c.cc.Invoke(ctx, UsersService_GetUserByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations should embed UnimplementedUsersServiceServer
// for forward compatibility.
type UsersServiceServer interface {
	SaveSecurityImage(context.Context, *SaveSecurityImageRequest) (*SaveSecurityImageResponse, error)
	GetSecurityImage(context.Context, *GetSecurityImageRequest) (*GetSecurityImageResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error)
}

// UnimplementedUsersServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServiceServer struct{}

func (UnimplementedUsersServiceServer) SaveSecurityImage(context.Context, *SaveSecurityImageRequest) (*SaveSecurityImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSecurityImage not implemented")
}
func (UnimplementedUsersServiceServer) GetSecurityImage(context.Context, *GetSecurityImageRequest) (*GetSecurityImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecurityImage not implemented")
}
func (UnimplementedUsersServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServiceServer) GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUsersServiceServer) testEmbeddedByValue() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	// If the following call pancis, it indicates UnimplementedUsersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_SaveSecurityImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveSecurityImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).SaveSecurityImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_SaveSecurityImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).SaveSecurityImage(ctx, req.(*SaveSecurityImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetSecurityImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecurityImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetSecurityImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_GetSecurityImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetSecurityImage(ctx, req.(*GetSecurityImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_GetUserByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUserByEmail(ctx, req.(*GetUserByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveSecurityImage",
			Handler:    _UsersService_SaveSecurityImage_Handler,
		},
		{
			MethodName: "GetSecurityImage",
			Handler:    _UsersService_GetSecurityImage_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UsersService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UsersService_GetUserByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
