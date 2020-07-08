package handler

import (
	"context"

	helloworld "../proto"
	"github.com/micro/go-micro/v2/util/log"
)

type Helloworld struct{}

func (g *Helloworld) Callme(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) {
	log.Log("Received Helloworld.Call request")
	rsp.Msg = "Hello" + req.Name
	return nil
}

func (g *Helloworld) Take(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) {
	log.Log("Received Helloworld.Call request")
	rsp.Msg = "Hello???" + req.Name
	return nil
}
