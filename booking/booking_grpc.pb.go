// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: booking.proto

package booking

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
	Booking_CreateBooking_FullMethodName         = "/booking_api.Booking/CreateBooking"
	Booking_GetBookingsByResource_FullMethodName = "/booking_api.Booking/GetBookingsByResource"
)

// BookingClient is the client API for Booking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	GetBookingsByResource(ctx context.Context, in *GetBookingsByResourceRequest, opts ...grpc.CallOption) (*GetBookingsByResourceResponse, error)
}

type bookingClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingClient(cc grpc.ClientConnInterface) BookingClient {
	return &bookingClient{cc}
}

func (c *bookingClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, Booking_CreateBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingClient) GetBookingsByResource(ctx context.Context, in *GetBookingsByResourceRequest, opts ...grpc.CallOption) (*GetBookingsByResourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookingsByResourceResponse)
	err := c.cc.Invoke(ctx, Booking_GetBookingsByResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServer is the server API for Booking service.
// All implementations must embed UnimplementedBookingServer
// for forward compatibility.
type BookingServer interface {
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	GetBookingsByResource(context.Context, *GetBookingsByResourceRequest) (*GetBookingsByResourceResponse, error)
	mustEmbedUnimplementedBookingServer()
}

// UnimplementedBookingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookingServer struct{}

func (UnimplementedBookingServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServer) GetBookingsByResource(context.Context, *GetBookingsByResourceRequest) (*GetBookingsByResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingsByResource not implemented")
}
func (UnimplementedBookingServer) mustEmbedUnimplementedBookingServer() {}
func (UnimplementedBookingServer) testEmbeddedByValue()                 {}

// UnsafeBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServer will
// result in compilation errors.
type UnsafeBookingServer interface {
	mustEmbedUnimplementedBookingServer()
}

func RegisterBookingServer(s grpc.ServiceRegistrar, srv BookingServer) {
	// If the following call pancis, it indicates UnimplementedBookingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Booking_ServiceDesc, srv)
}

func _Booking_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booking_CreateBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Booking_GetBookingsByResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingsByResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).GetBookingsByResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booking_GetBookingsByResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).GetBookingsByResource(ctx, req.(*GetBookingsByResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Booking_ServiceDesc is the grpc.ServiceDesc for Booking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Booking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking_api.Booking",
	HandlerType: (*BookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _Booking_CreateBooking_Handler,
		},
		{
			MethodName: "GetBookingsByResource",
			Handler:    _Booking_GetBookingsByResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
