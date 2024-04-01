// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/payment/payment.proto

package payment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	sdk "admin-gateway/proto/sdk"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	CreatePaymentIntent(ctx context.Context, in *MsgCreatePaymentIntent, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	CheckoutSession(ctx context.Context, in *MsgId, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	HookPayment(ctx context.Context, in *MsgPaymentIntent, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreatePaymentIntent(ctx context.Context, in *MsgCreatePaymentIntent, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/paymentService.paymentService/CreatePaymentIntent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CheckoutSession(ctx context.Context, in *MsgId, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/paymentService.paymentService/CheckoutSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) HookPayment(ctx context.Context, in *MsgPaymentIntent, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/paymentService.paymentService/HookPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	CreatePaymentIntent(context.Context, *MsgCreatePaymentIntent) (*sdk.BaseResponse, error)
	CheckoutSession(context.Context, *MsgId) (*sdk.BaseResponse, error)
	HookPayment(context.Context, *MsgPaymentIntent) (*sdk.BaseResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) CreatePaymentIntent(context.Context, *MsgCreatePaymentIntent) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePaymentIntent not implemented")
}
func (UnimplementedPaymentServiceServer) CheckoutSession(context.Context, *MsgId) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckoutSession not implemented")
}
func (UnimplementedPaymentServiceServer) HookPayment(context.Context, *MsgPaymentIntent) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HookPayment not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_CreatePaymentIntent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePaymentIntent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePaymentIntent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentService.paymentService/CreatePaymentIntent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePaymentIntent(ctx, req.(*MsgCreatePaymentIntent))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CheckoutSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CheckoutSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentService.paymentService/CheckoutSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CheckoutSession(ctx, req.(*MsgId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_HookPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPaymentIntent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).HookPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentService.paymentService/HookPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).HookPayment(ctx, req.(*MsgPaymentIntent))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paymentService.paymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePaymentIntent",
			Handler:    _PaymentService_CreatePaymentIntent_Handler,
		},
		{
			MethodName: "CheckoutSession",
			Handler:    _PaymentService_CheckoutSession_Handler,
		},
		{
			MethodName: "HookPayment",
			Handler:    _PaymentService_HookPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/payment/payment.proto",
}