// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bookgateway.proto

package bookgateway

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Api Endpoints for Bookgateway service

func NewBookgatewayEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Bookgateway service

type BookgatewayService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*Response, error)
}

type bookgatewayService struct {
	c    client.Client
	name string
}

func NewBookgatewayService(name string, c client.Client) BookgatewayService {
	return &bookgatewayService{
		c:    c,
		name: name,
	}
}

func (c *bookgatewayService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Bookgateway.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookgatewayService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Bookgateway.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bookgateway service

type BookgatewayHandler interface {
	Call(context.Context, *Request, *Response) error
	List(context.Context, *ListRequest, *Response) error
}

func RegisterBookgatewayHandler(s server.Server, hdlr BookgatewayHandler, opts ...server.HandlerOption) error {
	type bookgateway interface {
		Call(ctx context.Context, in *Request, out *Response) error
		List(ctx context.Context, in *ListRequest, out *Response) error
	}
	type Bookgateway struct {
		bookgateway
	}
	h := &bookgatewayHandler{hdlr}
	return s.Handle(s.NewHandler(&Bookgateway{h}, opts...))
}

type bookgatewayHandler struct {
	BookgatewayHandler
}

func (h *bookgatewayHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.BookgatewayHandler.Call(ctx, in, out)
}

func (h *bookgatewayHandler) List(ctx context.Context, in *ListRequest, out *Response) error {
	return h.BookgatewayHandler.List(ctx, in, out)
}
