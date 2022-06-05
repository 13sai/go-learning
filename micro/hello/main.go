package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/13sai/go-learing/micro/hello/common"
	"github.com/13sai/go-learing/micro/hello/hello"
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
	"github.com/sirupsen/logrus"
)

const (
	ServiceName = "hello-server"
)

type HelloServer struct{}

func (s *HelloServer) SayHello(ctx context.Context, req *hello.HelloRequest, res *hello.HelloReply) error {
	res.Message = "hello " + req.Name
	res.Hello = common.TypeHello_Afernoon
	return nil
}

func main() {
	service := micro.NewService(
		// Set service name
		micro.Name(ServiceName),
		micro.AfterStart(func() error {
			fmt.Println("starting...")
			return nil
		}),
		micro.Address(":8089"),
	)

	service.Init()

	hello.RegisterDemoHandler(service.Server(), &HelloServer{})

	go func() {
		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go func() {
		tick := time.NewTicker(3 * time.Second)

		for {
			select {
			case <-stop:
				tick.Stop()
			default:
				<-tick.C
				client()
			}
		}
	}()

	select {
	case <-stop:
		logrus.Infof("got exit signal, shutdown")
	}
}

func client() {
	service := micro.NewService(micro.Name(ServiceName + "client"))
	c := hello.NewDemoService(ServiceName, service.Client())

	// 发起RPC调用
	rsp, err := c.SayHello(context.TODO(), &hello.HelloRequest{Name: "13sai"})
	if err != nil {
		fmt.Println(err)
	}

	// 打印返回值
	fmt.Println(rsp.Message)
}
