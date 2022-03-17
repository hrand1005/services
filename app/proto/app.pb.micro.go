// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/app.proto

package app

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

// Api Endpoints for App service

func NewAppEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for App service

type AppService interface {
	Reserve(ctx context.Context, in *ReserveRequest, opts ...client.CallOption) (*ReserveResponse, error)
	Regions(ctx context.Context, in *RegionsRequest, opts ...client.CallOption) (*RegionsResponse, error)
	Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Resolve(ctx context.Context, in *ResolveRequest, opts ...client.CallOption) (*ResolveResponse, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...client.CallOption) (*LogsResponse, error)
}

type appService struct {
	c    client.Client
	name string
}

func NewAppService(name string, c client.Client) AppService {
	return &appService{
		c:    c,
		name: name,
	}
}

func (c *appService) Reserve(ctx context.Context, in *ReserveRequest, opts ...client.CallOption) (*ReserveResponse, error) {
	req := c.c.NewRequest(c.name, "App.Reserve", in)
	out := new(ReserveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) Regions(ctx context.Context, in *RegionsRequest, opts ...client.CallOption) (*RegionsResponse, error) {
	req := c.c.NewRequest(c.name, "App.Regions", in)
	out := new(RegionsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error) {
	req := c.c.NewRequest(c.name, "App.Run", in)
	out := new(RunResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "App.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "App.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error) {
	req := c.c.NewRequest(c.name, "App.Status", in)
	out := new(StatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "App.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) Resolve(ctx context.Context, in *ResolveRequest, opts ...client.CallOption) (*ResolveResponse, error) {
	req := c.c.NewRequest(c.name, "App.Resolve", in)
	out := new(ResolveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) Logs(ctx context.Context, in *LogsRequest, opts ...client.CallOption) (*LogsResponse, error) {
	req := c.c.NewRequest(c.name, "App.Logs", in)
	out := new(LogsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for App service

type AppHandler interface {
	Reserve(context.Context, *ReserveRequest, *ReserveResponse) error
	Regions(context.Context, *RegionsRequest, *RegionsResponse) error
	Run(context.Context, *RunRequest, *RunResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Status(context.Context, *StatusRequest, *StatusResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Resolve(context.Context, *ResolveRequest, *ResolveResponse) error
	Logs(context.Context, *LogsRequest, *LogsResponse) error
}

func RegisterAppHandler(s server.Server, hdlr AppHandler, opts ...server.HandlerOption) error {
	type app interface {
		Reserve(ctx context.Context, in *ReserveRequest, out *ReserveResponse) error
		Regions(ctx context.Context, in *RegionsRequest, out *RegionsResponse) error
		Run(ctx context.Context, in *RunRequest, out *RunResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Status(ctx context.Context, in *StatusRequest, out *StatusResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Resolve(ctx context.Context, in *ResolveRequest, out *ResolveResponse) error
		Logs(ctx context.Context, in *LogsRequest, out *LogsResponse) error
	}
	type App struct {
		app
	}
	h := &appHandler{hdlr}
	return s.Handle(s.NewHandler(&App{h}, opts...))
}

type appHandler struct {
	AppHandler
}

func (h *appHandler) Reserve(ctx context.Context, in *ReserveRequest, out *ReserveResponse) error {
	return h.AppHandler.Reserve(ctx, in, out)
}

func (h *appHandler) Regions(ctx context.Context, in *RegionsRequest, out *RegionsResponse) error {
	return h.AppHandler.Regions(ctx, in, out)
}

func (h *appHandler) Run(ctx context.Context, in *RunRequest, out *RunResponse) error {
	return h.AppHandler.Run(ctx, in, out)
}

func (h *appHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.AppHandler.Update(ctx, in, out)
}

func (h *appHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.AppHandler.Delete(ctx, in, out)
}

func (h *appHandler) Status(ctx context.Context, in *StatusRequest, out *StatusResponse) error {
	return h.AppHandler.Status(ctx, in, out)
}

func (h *appHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.AppHandler.List(ctx, in, out)
}

func (h *appHandler) Resolve(ctx context.Context, in *ResolveRequest, out *ResolveResponse) error {
	return h.AppHandler.Resolve(ctx, in, out)
}

func (h *appHandler) Logs(ctx context.Context, in *LogsRequest, out *LogsResponse) error {
	return h.AppHandler.Logs(ctx, in, out)
}
