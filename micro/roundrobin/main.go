package main

import (
	"context"
	"fmt"

	"github.com/13sai/go-learing/micro/hello/hello"
	"github.com/micro/go-micro/v2"
	roundrobin "github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
)
type HelloServer struct{}

func (s *HelloServer) SayHello(ctx context.Context, req *hello.HelloRequest, res *hello.HelloReply) error {
	res.Message = "hello " + req.Name
	return nil
}

func main() {
	wrapper := roundrobin.NewClientWrapper()
	service := micro.NewService(
		micro.Name("roundrobin"),
		micro.Address(":8081"),
		micro.AfterStart(func() error {
			fmt.Println("start successful!")
			return nil
		}),
		micro.WrapClient(wrapper),
	)

	service.Init()

	// service.Server().Handle(
	// 	service.Server().NewHandler(
	// 		&HelloServer{Client: hello.NewDemoService("go.micro.demo", service.Client())},
	// 	),
	// )
	hello.RegisterDemoHandler(service.Server(), &HelloServer{})

	service.Run()
}

