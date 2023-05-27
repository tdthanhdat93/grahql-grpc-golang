// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.4
// source: flight.proto

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

// FlightServiceClient is the client API for FlightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlightServiceClient interface {
	CreateFlight(ctx context.Context, in *FlightInfo, opts ...grpc.CallOption) (*FlightInfo, error)
	UpdateFlight(ctx context.Context, in *FlightInfo, opts ...grpc.CallOption) (*FlightInfo, error)
	ListFlights(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (FlightService_ListFlightsClient, error)
	FindFlightsByPlaneId(ctx context.Context, in *FindFlightByPlaneIdRequest, opts ...grpc.CallOption) (FlightService_FindFlightsByPlaneIdClient, error)
	FindFlights(ctx context.Context, in *FindFlightRequest, opts ...grpc.CallOption) (FlightService_FindFlightsClient, error)
	BookFlight(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*FlightInfo, error)
	DeleteFlight(ctx context.Context, in *DeleteFlightByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type flightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightServiceClient(cc grpc.ClientConnInterface) FlightServiceClient {
	return &flightServiceClient{cc}
}

func (c *flightServiceClient) CreateFlight(ctx context.Context, in *FlightInfo, opts ...grpc.CallOption) (*FlightInfo, error) {
	out := new(FlightInfo)
	err := c.cc.Invoke(ctx, "/proto.FlightService/CreateFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) UpdateFlight(ctx context.Context, in *FlightInfo, opts ...grpc.CallOption) (*FlightInfo, error) {
	out := new(FlightInfo)
	err := c.cc.Invoke(ctx, "/proto.FlightService/UpdateFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) ListFlights(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (FlightService_ListFlightsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlightService_ServiceDesc.Streams[0], "/proto.FlightService/ListFlights", opts...)
	if err != nil {
		return nil, err
	}
	x := &flightServiceListFlightsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlightService_ListFlightsClient interface {
	Recv() (*FlightInfo, error)
	grpc.ClientStream
}

type flightServiceListFlightsClient struct {
	grpc.ClientStream
}

func (x *flightServiceListFlightsClient) Recv() (*FlightInfo, error) {
	m := new(FlightInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flightServiceClient) FindFlightsByPlaneId(ctx context.Context, in *FindFlightByPlaneIdRequest, opts ...grpc.CallOption) (FlightService_FindFlightsByPlaneIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlightService_ServiceDesc.Streams[1], "/proto.FlightService/FindFlightsByPlaneId", opts...)
	if err != nil {
		return nil, err
	}
	x := &flightServiceFindFlightsByPlaneIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlightService_FindFlightsByPlaneIdClient interface {
	Recv() (*FlightInfo, error)
	grpc.ClientStream
}

type flightServiceFindFlightsByPlaneIdClient struct {
	grpc.ClientStream
}

func (x *flightServiceFindFlightsByPlaneIdClient) Recv() (*FlightInfo, error) {
	m := new(FlightInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flightServiceClient) FindFlights(ctx context.Context, in *FindFlightRequest, opts ...grpc.CallOption) (FlightService_FindFlightsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlightService_ServiceDesc.Streams[2], "/proto.FlightService/FindFlights", opts...)
	if err != nil {
		return nil, err
	}
	x := &flightServiceFindFlightsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlightService_FindFlightsClient interface {
	Recv() (*FlightInfo, error)
	grpc.ClientStream
}

type flightServiceFindFlightsClient struct {
	grpc.ClientStream
}

func (x *flightServiceFindFlightsClient) Recv() (*FlightInfo, error) {
	m := new(FlightInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flightServiceClient) BookFlight(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*FlightInfo, error) {
	out := new(FlightInfo)
	err := c.cc.Invoke(ctx, "/proto.FlightService/BookFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) DeleteFlight(ctx context.Context, in *DeleteFlightByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.FlightService/DeleteFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightServiceServer is the server API for FlightService service.
// All implementations must embed UnimplementedFlightServiceServer
// for forward compatibility
type FlightServiceServer interface {
	CreateFlight(context.Context, *FlightInfo) (*FlightInfo, error)
	UpdateFlight(context.Context, *FlightInfo) (*FlightInfo, error)
	ListFlights(*emptypb.Empty, FlightService_ListFlightsServer) error
	FindFlightsByPlaneId(*FindFlightByPlaneIdRequest, FlightService_FindFlightsByPlaneIdServer) error
	FindFlights(*FindFlightRequest, FlightService_FindFlightsServer) error
	BookFlight(context.Context, *BookingRequest) (*FlightInfo, error)
	DeleteFlight(context.Context, *DeleteFlightByIdRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedFlightServiceServer()
}

// UnimplementedFlightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlightServiceServer struct {
}

func (UnimplementedFlightServiceServer) CreateFlight(context.Context, *FlightInfo) (*FlightInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlight not implemented")
}
func (UnimplementedFlightServiceServer) UpdateFlight(context.Context, *FlightInfo) (*FlightInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlight not implemented")
}
func (UnimplementedFlightServiceServer) ListFlights(*emptypb.Empty, FlightService_ListFlightsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFlights not implemented")
}
func (UnimplementedFlightServiceServer) FindFlightsByPlaneId(*FindFlightByPlaneIdRequest, FlightService_FindFlightsByPlaneIdServer) error {
	return status.Errorf(codes.Unimplemented, "method FindFlightsByPlaneId not implemented")
}
func (UnimplementedFlightServiceServer) FindFlights(*FindFlightRequest, FlightService_FindFlightsServer) error {
	return status.Errorf(codes.Unimplemented, "method FindFlights not implemented")
}
func (UnimplementedFlightServiceServer) BookFlight(context.Context, *BookingRequest) (*FlightInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookFlight not implemented")
}
func (UnimplementedFlightServiceServer) DeleteFlight(context.Context, *DeleteFlightByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFlight not implemented")
}
func (UnimplementedFlightServiceServer) mustEmbedUnimplementedFlightServiceServer() {}

// UnsafeFlightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightServiceServer will
// result in compilation errors.
type UnsafeFlightServiceServer interface {
	mustEmbedUnimplementedFlightServiceServer()
}

func RegisterFlightServiceServer(s grpc.ServiceRegistrar, srv FlightServiceServer) {
	s.RegisterService(&FlightService_ServiceDesc, srv)
}

func _FlightService_CreateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).CreateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FlightService/CreateFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).CreateFlight(ctx, req.(*FlightInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_UpdateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).UpdateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FlightService/UpdateFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).UpdateFlight(ctx, req.(*FlightInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_ListFlights_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlightServiceServer).ListFlights(m, &flightServiceListFlightsServer{stream})
}

type FlightService_ListFlightsServer interface {
	Send(*FlightInfo) error
	grpc.ServerStream
}

type flightServiceListFlightsServer struct {
	grpc.ServerStream
}

func (x *flightServiceListFlightsServer) Send(m *FlightInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _FlightService_FindFlightsByPlaneId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindFlightByPlaneIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlightServiceServer).FindFlightsByPlaneId(m, &flightServiceFindFlightsByPlaneIdServer{stream})
}

type FlightService_FindFlightsByPlaneIdServer interface {
	Send(*FlightInfo) error
	grpc.ServerStream
}

type flightServiceFindFlightsByPlaneIdServer struct {
	grpc.ServerStream
}

func (x *flightServiceFindFlightsByPlaneIdServer) Send(m *FlightInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _FlightService_FindFlights_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindFlightRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlightServiceServer).FindFlights(m, &flightServiceFindFlightsServer{stream})
}

type FlightService_FindFlightsServer interface {
	Send(*FlightInfo) error
	grpc.ServerStream
}

type flightServiceFindFlightsServer struct {
	grpc.ServerStream
}

func (x *flightServiceFindFlightsServer) Send(m *FlightInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _FlightService_BookFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).BookFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FlightService/BookFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).BookFlight(ctx, req.(*BookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_DeleteFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFlightByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).DeleteFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FlightService/DeleteFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).DeleteFlight(ctx, req.(*DeleteFlightByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlightService_ServiceDesc is the grpc.ServiceDesc for FlightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FlightService",
	HandlerType: (*FlightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFlight",
			Handler:    _FlightService_CreateFlight_Handler,
		},
		{
			MethodName: "UpdateFlight",
			Handler:    _FlightService_UpdateFlight_Handler,
		},
		{
			MethodName: "BookFlight",
			Handler:    _FlightService_BookFlight_Handler,
		},
		{
			MethodName: "DeleteFlight",
			Handler:    _FlightService_DeleteFlight_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFlights",
			Handler:       _FlightService_ListFlights_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindFlightsByPlaneId",
			Handler:       _FlightService_FindFlightsByPlaneId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindFlights",
			Handler:       _FlightService_FindFlights_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "flight.proto",
}
