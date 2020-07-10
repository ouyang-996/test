package sayhandler

import (
	"context"
	"fmt"
	say "say-service/proto/msg"
	"say-service/service"
)

//SayHandler is a example
type SayHandler struct {
	say service.ISay
}

//NewSayHandler SayHandler example
func NewSayHandler() *SayHandler {
	return &SayHandler{
		say: service.NewSay(),
	}
}

func (s *SayHandler) Hello(ctx context.Context, req *say.HelloRequest, rsp *say.HelloResponse) error {
	msg := s.say.GetHello()
	fmt.Println("ddd")
	rsp.Data = msg + " " + req.Name
	return nil
}

func (s *SayHandler) Monming(ctx context.Context, req *say.MonmingRequest, rsp *say.MonmingResponse) error {
	msg := s.say.GetMonming()
	rsp.Data = msg + " " + req.Name
	return nil
}
