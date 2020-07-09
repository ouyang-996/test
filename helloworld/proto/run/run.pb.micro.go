// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: run.proto

package run

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Run service

type RunService interface {
	Gorun(ctx context.Context, in *Runrequest, opts ...client.CallOption) (*Runresponse, error)
}

type runService struct {
	c    client.Client
	name string
}

func NewRunService(name string, c client.Client) RunService {
	return &runService{
		c:    c,
		name: name,
	}
}

func (c *runService) Gorun(ctx context.Context, in *Runrequest, opts ...client.CallOption) (*Runresponse, error) {
	req := c.c.NewRequest(c.name, "Run.Gorun", in)
	out := new(Runresponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Run service

type RunHandler interface {
	Gorun(context.Context, *Runrequest, *Runresponse) error
}

func RegisterRunHandler(s server.Server, hdlr RunHandler, opts ...server.HandlerOption) error {
	type run interface {
		Gorun(ctx context.Context, in *Runrequest, out *Runresponse) error
	}
	type Run struct {
		run
	}
	h := &runHandler{hdlr}
	return s.Handle(s.NewHandler(&Run{h}, opts...))
}

type runHandler struct {
	RunHandler
}

func (h *runHandler) Gorun(ctx context.Context, in *Runrequest, out *Runresponse) error {
	return h.RunHandler.Gorun(ctx, in, out)
}
