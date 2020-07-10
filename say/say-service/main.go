package main

import (
	"fmt"

	"github.com/micro/go-micro/v2"

	msg "say-service/handler/say"

	say "say-service/proto/msg"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.say"),
	)

	service.Init()

	say.RegisterMsgServiceHandler(service.Server(), msg.NewSayHandler())

	if err := service.Run(); err != nil {
		fmt.Println("run?")
	}

}
