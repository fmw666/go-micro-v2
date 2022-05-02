// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderService.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/Allenxuxu/mMicro/api"
	client "github.com/aiscrm/go-micro/v2/client"
	server "github.com/aiscrm/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for OrderService service

func NewOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderService service

type OrderService interface {
	// rpc GetOrderList() returns(UserDetailResponse);
	GetOrderDetail(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderDetailResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) GetOrderDetail(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrderDetail", in)
	out := new(OrderDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	// rpc GetOrderList() returns(UserDetailResponse);
	GetOrderDetail(context.Context, *OrderRequest, *OrderDetailResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		GetOrderDetail(ctx context.Context, in *OrderRequest, out *OrderDetailResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) GetOrderDetail(ctx context.Context, in *OrderRequest, out *OrderDetailResponse) error {
	return h.OrderServiceHandler.GetOrderDetail(ctx, in, out)
}
