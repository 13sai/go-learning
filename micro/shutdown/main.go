package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// shutdown after 5 seconds
	go func() {
		<-time.After(time.Second * 5)
		log.Println("Shutdown example: shutting down service")
		cancel()
	}()

	service := micro.NewService(
		micro.Context(ctx),
		micro.BeforeStop(func () error {
			fmt.Println("shutdown...")
			return nil
		}),
		micro.AfterStop(func () error {
			fmt.Println("shutdown successful")
			return nil
		}),
	)

	// init service
	service.Init()

	// run service
	service.Run()
}