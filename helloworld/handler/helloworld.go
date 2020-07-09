package handler

import (
	"context"

	helloworld "helloworld/proto/helloworld"

	"github.com/micro/go-micro/v2/util/log"
)

type Helloworld struct{}

func (g *Helloworld) Callme(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	log.Log("Received Helloworld.Call request")

	method := req.GetName()
	// if method != "GET" {
	// 	rsp.Msg = "GOOD "
	// 	return nil
	// }
	// name := req.Get["name"].GetValues()
	rsp.Msg = "GOOD " + method
	return nil
}

func (g *Helloworld) Take(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	log.Log("Received Helloworld.Call request")
	rsp.Msg = "Hello???" + req.Name
	return nil
}
