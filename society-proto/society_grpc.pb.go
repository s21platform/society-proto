// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: society.proto

package society_proto

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
	SocietyService_CreateSociety_FullMethodName  = "/SocietyService/CreateSociety"
	SocietyService_GetAccessLevel_FullMethodName = "/SocietyService/GetAccessLevel"
	SocietyService_GetPermissions_FullMethodName = "/SocietyService/GetPermissions"
)

// SocietyServiceClient is the client API for SocietyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocietyServiceClient interface {
	CreateSociety(ctx context.Context, in *SetSocietyIn, opts ...grpc.CallOption) (*SetSocietyOut, error)
	GetAccessLevel(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAccessLevelOut, error)
	GetPermissions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPermissionsOut, error)
}

type societyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSocietyServiceClient(cc grpc.ClientConnInterface) SocietyServiceClient {
	return &societyServiceClient{cc}
}

func (c *societyServiceClient) CreateSociety(ctx context.Context, in *SetSocietyIn, opts ...grpc.CallOption) (*SetSocietyOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetSocietyOut)
	err := c.cc.Invoke(ctx, SocietyService_CreateSociety_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *societyServiceClient) GetAccessLevel(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAccessLevelOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAccessLevelOut)
	err := c.cc.Invoke(ctx, SocietyService_GetAccessLevel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *societyServiceClient) GetPermissions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPermissionsOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPermissionsOut)
	err := c.cc.Invoke(ctx, SocietyService_GetPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocietyServiceServer is the server API for SocietyService service.
// All implementations must embed UnimplementedSocietyServiceServer
// for forward compatibility.
type SocietyServiceServer interface {
	CreateSociety(context.Context, *SetSocietyIn) (*SetSocietyOut, error)
	GetAccessLevel(context.Context, *Empty) (*GetAccessLevelOut, error)
	GetPermissions(context.Context, *Empty) (*GetPermissionsOut, error)
	mustEmbedUnimplementedSocietyServiceServer()
}

// UnimplementedSocietyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSocietyServiceServer struct{}

func (UnimplementedSocietyServiceServer) CreateSociety(context.Context, *SetSocietyIn) (*SetSocietyOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSociety not implemented")
}
func (UnimplementedSocietyServiceServer) GetAccessLevel(context.Context, *Empty) (*GetAccessLevelOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessLevel not implemented")
}
func (UnimplementedSocietyServiceServer) GetPermissions(context.Context, *Empty) (*GetPermissionsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedSocietyServiceServer) mustEmbedUnimplementedSocietyServiceServer() {}
func (UnimplementedSocietyServiceServer) testEmbeddedByValue()                        {}

// UnsafeSocietyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocietyServiceServer will
// result in compilation errors.
type UnsafeSocietyServiceServer interface {
	mustEmbedUnimplementedSocietyServiceServer()
}

func RegisterSocietyServiceServer(s grpc.ServiceRegistrar, srv SocietyServiceServer) {
	// If the following call pancis, it indicates UnimplementedSocietyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SocietyService_ServiceDesc, srv)
}

func _SocietyService_CreateSociety_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSocietyIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).CreateSociety(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_CreateSociety_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).CreateSociety(ctx, req.(*SetSocietyIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocietyService_GetAccessLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).GetAccessLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_GetAccessLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).GetAccessLevel(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocietyService_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_GetPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).GetPermissions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SocietyService_ServiceDesc is the grpc.ServiceDesc for SocietyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SocietyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SocietyService",
	HandlerType: (*SocietyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSociety",
			Handler:    _SocietyService_CreateSociety_Handler,
		},
		{
			MethodName: "GetAccessLevel",
			Handler:    _SocietyService_GetAccessLevel_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _SocietyService_GetPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "society.proto",
}
