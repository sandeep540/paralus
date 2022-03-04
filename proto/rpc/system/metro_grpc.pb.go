// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpcv3

import (
	context "context"
	v3 "github.com/RafaySystems/rcloud-base/proto/types/infrapb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LocationClient is the client API for Location service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationClient interface {
	CreateLocation(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.Location, error)
	GetLocations(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.LocationList, error)
	GetLocation(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.Location, error)
	UpdateLocation(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.Location, error)
	DeleteLocation(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.Location, error)
}

type locationClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationClient(cc grpc.ClientConnInterface) LocationClient {
	return &locationClient{cc}
}

func (c *locationClient) CreateLocation(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.Location, error) {
	out := new(v3.Location)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Location/CreateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) GetLocations(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.LocationList, error) {
	out := new(v3.LocationList)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Location/GetLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) GetLocation(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.Location, error) {
	out := new(v3.Location)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Location/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) UpdateLocation(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.Location, error) {
	out := new(v3.Location)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Location/UpdateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) DeleteLocation(ctx context.Context, in *v3.Location, opts ...grpc.CallOption) (*v3.Location, error) {
	out := new(v3.Location)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Location/DeleteLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationServer is the server API for Location service.
// All implementations should embed UnimplementedLocationServer
// for forward compatibility
type LocationServer interface {
	CreateLocation(context.Context, *v3.Location) (*v3.Location, error)
	GetLocations(context.Context, *v3.Location) (*v3.LocationList, error)
	GetLocation(context.Context, *v3.Location) (*v3.Location, error)
	UpdateLocation(context.Context, *v3.Location) (*v3.Location, error)
	DeleteLocation(context.Context, *v3.Location) (*v3.Location, error)
}

// UnimplementedLocationServer should be embedded to have forward compatible implementations.
type UnimplementedLocationServer struct {
}

func (UnimplementedLocationServer) CreateLocation(context.Context, *v3.Location) (*v3.Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocation not implemented")
}
func (UnimplementedLocationServer) GetLocations(context.Context, *v3.Location) (*v3.LocationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocations not implemented")
}
func (UnimplementedLocationServer) GetLocation(context.Context, *v3.Location) (*v3.Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedLocationServer) UpdateLocation(context.Context, *v3.Location) (*v3.Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedLocationServer) DeleteLocation(context.Context, *v3.Location) (*v3.Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocation not implemented")
}

// UnsafeLocationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServer will
// result in compilation errors.
type UnsafeLocationServer interface {
	mustEmbedUnimplementedLocationServer()
}

func RegisterLocationServer(s grpc.ServiceRegistrar, srv LocationServer) {
	s.RegisterService(&Location_ServiceDesc, srv)
}

func _Location_CreateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).CreateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Location/CreateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).CreateLocation(ctx, req.(*v3.Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Location_GetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).GetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Location/GetLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).GetLocations(ctx, req.(*v3.Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Location_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Location/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).GetLocation(ctx, req.(*v3.Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Location_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Location/UpdateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).UpdateLocation(ctx, req.(*v3.Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Location_DeleteLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).DeleteLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Location/DeleteLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).DeleteLocation(ctx, req.(*v3.Location))
	}
	return interceptor(ctx, in, info, handler)
}

// Location_ServiceDesc is the grpc.ServiceDesc for Location service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Location_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rafay.dev.rpc.v3.Location",
	HandlerType: (*LocationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLocation",
			Handler:    _Location_CreateLocation_Handler,
		},
		{
			MethodName: "GetLocations",
			Handler:    _Location_GetLocations_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _Location_GetLocation_Handler,
		},
		{
			MethodName: "UpdateLocation",
			Handler:    _Location_UpdateLocation_Handler,
		},
		{
			MethodName: "DeleteLocation",
			Handler:    _Location_DeleteLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/system/metro.proto",
}
