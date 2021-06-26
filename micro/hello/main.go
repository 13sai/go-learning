package main

import (
	"context"
	"fmt"

	"github.com/13sai/go-learing/micro/hello/hello"
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
)

const (
	ServiceName = "hello-server"
)

type HelloServer struct{}

func (s *HelloServer) SayHello(ctx context.Context, req *hello.HelloRequest, res *hello.HelloReply) error {
	res.Message = "hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		// Set service name
		micro.Name(ServiceName),
		micro.BeforeStart(func() error {
			fmt.Println("starting...")
			return nil
		}),
	)

	service.Init()

	hello.RegisterDemoHandler(service.Server(), &HelloServer{})

	service.Run()
}

