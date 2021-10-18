// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authpb

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

// SignupServiceClient is the client API for SignupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignupServiceClient interface {
	SignUp(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	IsUsernameAvailable(ctx context.Context, in *UsernameAvailabilityRequest, opts ...grpc.CallOption) (*AvailabilityResponse, error)
	IsEmailAvailable(ctx context.Context, in *EmailAvailabilityRequest, opts ...grpc.CallOption) (*AvailabilityResponse, error)
}

type signupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignupServiceClient(cc grpc.ClientConnInterface) SignupServiceClient {
	return &signupServiceClient{cc}
}

func (c *signupServiceClient) SignUp(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/auth.SignupService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) IsUsernameAvailable(ctx context.Context, in *UsernameAvailabilityRequest, opts ...grpc.CallOption) (*AvailabilityResponse, error) {
	out := new(AvailabilityResponse)
	err := c.cc.Invoke(ctx, "/auth.SignupService/IsUsernameAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) IsEmailAvailable(ctx context.Context, in *EmailAvailabilityRequest, opts ...grpc.CallOption) (*AvailabilityResponse, error) {
	out := new(AvailabilityResponse)
	err := c.cc.Invoke(ctx, "/auth.SignupService/IsEmailAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupServiceServer is the server API for SignupService service.
// All implementations must embed UnimplementedSignupServiceServer
// for forward compatibility
type SignupServiceServer interface {
	SignUp(context.Context, *SignupRequest) (*SignupResponse, error)
	IsUsernameAvailable(context.Context, *UsernameAvailabilityRequest) (*AvailabilityResponse, error)
	IsEmailAvailable(context.Context, *EmailAvailabilityRequest) (*AvailabilityResponse, error)
	mustEmbedUnimplementedSignupServiceServer()
}

// UnimplementedSignupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSignupServiceServer struct {
}

func (UnimplementedSignupServiceServer) SignUp(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedSignupServiceServer) IsUsernameAvailable(context.Context, *UsernameAvailabilityRequest) (*AvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUsernameAvailable not implemented")
}
func (UnimplementedSignupServiceServer) IsEmailAvailable(context.Context, *EmailAvailabilityRequest) (*AvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsEmailAvailable not implemented")
}
func (UnimplementedSignupServiceServer) mustEmbedUnimplementedSignupServiceServer() {}

// UnsafeSignupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignupServiceServer will
// result in compilation errors.
type UnsafeSignupServiceServer interface {
	mustEmbedUnimplementedSignupServiceServer()
}

func RegisterSignupServiceServer(s grpc.ServiceRegistrar, srv SignupServiceServer) {
	s.RegisterService(&SignupService_ServiceDesc, srv)
}

func _SignupService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.SignupService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).SignUp(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_IsUsernameAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsernameAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).IsUsernameAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.SignupService/IsUsernameAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).IsUsernameAvailable(ctx, req.(*UsernameAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_IsEmailAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).IsEmailAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.SignupService/IsEmailAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).IsEmailAvailable(ctx, req.(*EmailAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignupService_ServiceDesc is the grpc.ServiceDesc for SignupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.SignupService",
	HandlerType: (*SignupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _SignupService_SignUp_Handler,
		},
		{
			MethodName: "IsUsernameAvailable",
			Handler:    _SignupService_IsUsernameAvailable_Handler,
		},
		{
			MethodName: "IsEmailAvailable",
			Handler:    _SignupService_IsEmailAvailable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
