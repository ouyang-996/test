package main

import (
	"fmt"
	"helloworld/handler"
	helloworld "helloworld/proto/helloworld"
	run "helloworld/proto/run"
	"time"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.helloworld"),
		micro.Version("latest"),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))
	run.RegisterRunHandler(service.Server(), new(handler.Run))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
