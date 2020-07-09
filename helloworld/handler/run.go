package handler

import (
	"context"
	run "helloworld/proto/run"
)

type Run struct{}

func (r *Run) Gorun(ctx context.Context, req *run.Runrequest, rsp *run.Runresponse) error {
	tag := req.GetTag()
	rsp.Msg = "I am " + tag
	return nil
}