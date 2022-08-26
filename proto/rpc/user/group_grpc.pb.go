// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/user/group.proto

package userv3

import (
	context "context"
	v31 "github.com/paralus/paralus/proto/types/commonpb/v3"
	v3 "github.com/paralus/paralus/proto/types/userpb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	CreateGroup(ctx context.Context, in *v3.Group, opts ...grpc.CallOption) (*v3.Group, error)
	GetGroups(ctx context.Context, in *v31.QueryOptions, opts ...grpc.CallOption) (*v3.GroupList, error)
	GetGroup(ctx context.Context, in *v3.Group, opts ...grpc.CallOption) (*v3.Group, error)
	UpdateGroup(ctx context.Context, in *v3.Group, opts ...grpc.CallOption) (*v3.Group, error)
	DeleteGroup(ctx context.Context, in *v3.Group, opts ...grpc.CallOption) (*v3.Group, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *v3.Group, opts ...grpc.CallOption) (*v3.Group, error) {
	out := new(v3.Group)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.GroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroups(ctx context.Context, in *v31.QueryOptions, opts ...grpc.CallOption) (*v3.GroupList, error) {
	out := new(v3.GroupList)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.GroupService/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroup(ctx context.Context, in *v3.Group, opts ...grpc.CallOption) (*v3.Group, error) {
	out := new(v3.Group)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.GroupService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *v3.Group, opts ...grpc.CallOption) (*v3.Group, error) {
	out := new(v3.Group)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.GroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *v3.Group, opts ...grpc.CallOption) (*v3.Group, error) {
	out := new(v3.Group)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.GroupService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations should embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	CreateGroup(context.Context, *v3.Group) (*v3.Group, error)
	GetGroups(context.Context, *v31.QueryOptions) (*v3.GroupList, error)
	GetGroup(context.Context, *v3.Group) (*v3.Group, error)
	UpdateGroup(context.Context, *v3.Group) (*v3.Group, error)
	DeleteGroup(context.Context, *v3.Group) (*v3.Group, error)
}

// UnimplementedGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) CreateGroup(context.Context, *v3.Group) (*v3.Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupServiceServer) GetGroups(context.Context, *v31.QueryOptions) (*v3.GroupList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedGroupServiceServer) GetGroup(context.Context, *v3.Group) (*v3.Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedGroupServiceServer) UpdateGroup(context.Context, *v3.Group) (*v3.Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedGroupServiceServer) DeleteGroup(context.Context, *v3.Group) (*v3.Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.GroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*v3.Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v31.QueryOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.GroupService/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroups(ctx, req.(*v31.QueryOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.GroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroup(ctx, req.(*v3.Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.GroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroup(ctx, req.(*v3.Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.GroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*v3.Group))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paralus.dev.rpc.user.v3.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _GroupService_GetGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _GroupService_GetGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/user/group.proto",
}
