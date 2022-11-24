// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: pb/order_pb.proto

package pb

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	CheckoutOrder(ctx context.Context, in *CheckoutOrderRequest, opts ...grpc.CallOption) (*CheckoutOrderReply, error)
	AddOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	GetOrderDetail(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Order, error)
	GetOrderList(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderListReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CheckoutOrder(ctx context.Context, in *CheckoutOrderRequest, opts ...grpc.CallOption) (*CheckoutOrderReply, error) {
	out := new(CheckoutOrderReply)
	err := c.cc.Invoke(ctx, "/pb.Order/CheckoutOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) AddOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/pb.Order/AddOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderDetail(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/pb.Order/GetOrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderList(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderListReply, error) {
	out := new(OrderListReply)
	err := c.cc.Invoke(ctx, "/pb.Order/GetOrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations should embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	CheckoutOrder(context.Context, *CheckoutOrderRequest) (*CheckoutOrderReply, error)
	AddOrder(context.Context, *Order) (*Order, error)
	GetOrderDetail(context.Context, *OrderRequest) (*Order, error)
	GetOrderList(context.Context, *OrderRequest) (*OrderListReply, error)
}

// UnimplementedOrderServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CheckoutOrder(context.Context, *CheckoutOrderRequest) (*CheckoutOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckoutOrder not implemented")
}
func (UnimplementedOrderServer) AddOrder(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (UnimplementedOrderServer) GetOrderDetail(context.Context, *OrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetail not implemented")
}
func (UnimplementedOrderServer) GetOrderList(context.Context, *OrderRequest) (*OrderListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CheckoutOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CheckoutOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Order/CheckoutOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CheckoutOrder(ctx, req.(*CheckoutOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Order/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).AddOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Order/GetOrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderDetail(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Order/GetOrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderList(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckoutOrder",
			Handler:    _Order_CheckoutOrder_Handler,
		},
		{
			MethodName: "AddOrder",
			Handler:    _Order_AddOrder_Handler,
		},
		{
			MethodName: "GetOrderDetail",
			Handler:    _Order_GetOrderDetail_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _Order_GetOrderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/order_pb.proto",
}
