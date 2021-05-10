// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hall.proto

package hall

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

// Api Endpoints for Registery service

func NewRegisteryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Registery service

type RegisteryService interface {
	Ping(ctx context.Context, in *RegisteryRequest, opts ...client.CallOption) (*RegisteryResponse, error)
}

type registeryService struct {
	c    client.Client
	name string
}

func NewRegisteryService(name string, c client.Client) RegisteryService {
	return &registeryService{
		c:    c,
		name: name,
	}
}

func (c *registeryService) Ping(ctx context.Context, in *RegisteryRequest, opts ...client.CallOption) (*RegisteryResponse, error) {
	req := c.c.NewRequest(c.name, "Registery.Ping", in)
	out := new(RegisteryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registery service

type RegisteryHandler interface {
	Ping(context.Context, *RegisteryRequest, *RegisteryResponse) error
}

func RegisterRegisteryHandler(s server.Server, hdlr RegisteryHandler, opts ...server.HandlerOption) error {
	type registery interface {
		Ping(ctx context.Context, in *RegisteryRequest, out *RegisteryResponse) error
	}
	type Registery struct {
		registery
	}
	h := &registeryHandler{hdlr}
	return s.Handle(s.NewHandler(&Registery{h}, opts...))
}

type registeryHandler struct {
	RegisteryHandler
}

func (h *registeryHandler) Ping(ctx context.Context, in *RegisteryRequest, out *RegisteryResponse) error {
	return h.RegisteryHandler.Ping(ctx, in, out)
}
