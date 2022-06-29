package main

import (
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	hystrixDo()
}

func hystrixDo() {
	command := "13sai"
	hystrix.ConfigureCommand(command, hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  10,
		RequestVolumeThreshold: 10,
		SleepWindow:            2000,
	})

	defer hystrix.Flush()

	for i := 0; i < 100; i++ {
		err := hystrix.Do(command, func() error {
			_, err := http.Get("https://www.baidu1.com")
			fmt.Println("https://www.baidu1.com")
			return err
		}, func(err error) error {
			fmt.Printf("err=%v", err)
			return nil
		})
		fmt.Printf("i=%d,err=%+v---%v", i, command, err)
		fmt.Println()
	}
}
