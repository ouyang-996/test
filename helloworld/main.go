package main

import (
	"fmt"

	"github.com/micro/examples/helloworld/handler"
	"github.com/micro/go-micro/v2"

	helloworld "github.com/micro/examples/helloworld/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
	)

	service.Init()
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
