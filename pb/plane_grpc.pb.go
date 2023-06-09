// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.4
// source: plane.proto

package pb

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

// PlaneServiceClient is the client API for PlaneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaneServiceClient interface {
	CreatePlane(ctx context.Context, in *PlaneInfo, opts ...grpc.CallOption) (*PlaneInfo, error)
	UpdatePlane(ctx context.Context, in *PlaneInfo, opts ...grpc.CallOption) (*PlaneInfo, error)
	ListPlanes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PlaneService_ListPlanesClient, error)
	FindPlane(ctx context.Context, in *FindPlaneByIdRequest, opts ...grpc.CallOption) (*PlaneInfo, error)
	DeletePlane(ctx context.Context, in *DeletePlaneByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type planeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaneServiceClient(cc grpc.ClientConnInterface) PlaneServiceClient {
	return &planeServiceClient{cc}
}

func (c *planeServiceClient) CreatePlane(ctx context.Context, in *PlaneInfo, opts ...grpc.CallOption) (*PlaneInfo, error) {
	out := new(PlaneInfo)
	err := c.cc.Invoke(ctx, "/proto.PlaneService/CreatePlane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planeServiceClient) UpdatePlane(ctx context.Context, in *PlaneInfo, opts ...grpc.CallOption) (*PlaneInfo, error) {
	out := new(PlaneInfo)
	err := c.cc.Invoke(ctx, "/proto.PlaneService/UpdatePlane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planeServiceClient) ListPlanes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PlaneService_ListPlanesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PlaneService_ServiceDesc.Streams[0], "/proto.PlaneService/ListPlanes", opts...)
	if err != nil {
		return nil, err
	}
	x := &planeServiceListPlanesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlaneService_ListPlanesClient interface {
	Recv() (*PlaneInfo, error)
	grpc.ClientStream
}

type planeServiceListPlanesClient struct {
	grpc.ClientStream
}

func (x *planeServiceListPlanesClient) Recv() (*PlaneInfo, error) {
	m := new(PlaneInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *planeServiceClient) FindPlane(ctx context.Context, in *FindPlaneByIdRequest, opts ...grpc.CallOption) (*PlaneInfo, error) {
	out := new(PlaneInfo)
	err := c.cc.Invoke(ctx, "/proto.PlaneService/FindPlane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planeServiceClient) DeletePlane(ctx context.Context, in *DeletePlaneByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.PlaneService/DeletePlane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaneServiceServer is the server API for PlaneService service.
// All implementations must embed UnimplementedPlaneServiceServer
// for forward compatibility
type PlaneServiceServer interface {
	CreatePlane(context.Context, *PlaneInfo) (*PlaneInfo, error)
	UpdatePlane(context.Context, *PlaneInfo) (*PlaneInfo, error)
	ListPlanes(*emptypb.Empty, PlaneService_ListPlanesServer) error
	FindPlane(context.Context, *FindPlaneByIdRequest) (*PlaneInfo, error)
	DeletePlane(context.Context, *DeletePlaneByIdRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPlaneServiceServer()
}

// UnimplementedPlaneServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlaneServiceServer struct {
}

func (UnimplementedPlaneServiceServer) CreatePlane(context.Context, *PlaneInfo) (*PlaneInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlane not implemented")
}
func (UnimplementedPlaneServiceServer) UpdatePlane(context.Context, *PlaneInfo) (*PlaneInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlane not implemented")
}
func (UnimplementedPlaneServiceServer) ListPlanes(*emptypb.Empty, PlaneService_ListPlanesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPlanes not implemented")
}
func (UnimplementedPlaneServiceServer) FindPlane(context.Context, *FindPlaneByIdRequest) (*PlaneInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPlane not implemented")
}
func (UnimplementedPlaneServiceServer) DeletePlane(context.Context, *DeletePlaneByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlane not implemented")
}
func (UnimplementedPlaneServiceServer) mustEmbedUnimplementedPlaneServiceServer() {}

// UnsafePlaneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaneServiceServer will
// result in compilation errors.
type UnsafePlaneServiceServer interface {
	mustEmbedUnimplementedPlaneServiceServer()
}

func RegisterPlaneServiceServer(s grpc.ServiceRegistrar, srv PlaneServiceServer) {
	s.RegisterService(&PlaneService_ServiceDesc, srv)
}

func _PlaneService_CreatePlane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaneInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaneServiceServer).CreatePlane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlaneService/CreatePlane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaneServiceServer).CreatePlane(ctx, req.(*PlaneInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaneService_UpdatePlane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaneInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaneServiceServer).UpdatePlane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlaneService/UpdatePlane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaneServiceServer).UpdatePlane(ctx, req.(*PlaneInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaneService_ListPlanes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlaneServiceServer).ListPlanes(m, &planeServiceListPlanesServer{stream})
}

type PlaneService_ListPlanesServer interface {
	Send(*PlaneInfo) error
	grpc.ServerStream
}

type planeServiceListPlanesServer struct {
	grpc.ServerStream
}

func (x *planeServiceListPlanesServer) Send(m *PlaneInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _PlaneService_FindPlane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPlaneByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaneServiceServer).FindPlane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlaneService/FindPlane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaneServiceServer).FindPlane(ctx, req.(*FindPlaneByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaneService_DeletePlane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlaneByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaneServiceServer).DeletePlane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlaneService/DeletePlane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaneServiceServer).DeletePlane(ctx, req.(*DeletePlaneByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlaneService_ServiceDesc is the grpc.ServiceDesc for PlaneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PlaneService",
	HandlerType: (*PlaneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlane",
			Handler:    _PlaneService_CreatePlane_Handler,
		},
		{
			MethodName: "UpdatePlane",
			Handler:    _PlaneService_UpdatePlane_Handler,
		},
		{
			MethodName: "FindPlane",
			Handler:    _PlaneService_FindPlane_Handler,
		},
		{
			MethodName: "DeletePlane",
			Handler:    _PlaneService_DeletePlane_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPlanes",
			Handler:       _PlaneService_ListPlanes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "plane.proto",
}
