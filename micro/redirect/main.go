package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	api "github.com/micro/go-micro/v2/api/proto"
)
type Redirect struct{}

func(r *Redirect) Url(ctx context.Context, req *api.Request, res *api.Response) error {
	res.StatusCode = int32(301)
	res.Header = map[string]*api.Pair{
		"Location": &api.Pair{
			Key:    "Location",
			Values: []string{"https://github.13sai.com"},
		},
	}
	return nil
}


func main() {
	service := micro.NewService(
		micro.Name("redirect"),
		micro.Address(":8082"),
		micro.AfterStart(func() error {
			fmt.Println("start successful!")
			return nil
		}),
	)

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(new(Redirect)),
	)

	service.Run()
}

