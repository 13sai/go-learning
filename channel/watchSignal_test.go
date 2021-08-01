package channel

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	shut := make(chan os.Signal)
	// stopChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		watchCtx(ctx)
	}()
	// go func() {
	// 	watchChan(stopChan)
	// }()
	signal.Notify(shut, os.Interrupt)
	fmt.Println("开始监听信号")
	<-shut
	fmt.Println("中断，发起协程中断信号")
	cancel()
	// stopChan <- true
	// time.Sleep(1 * time.Second)
	fmt.Println("进程结束")
}

// 使用无缓冲通道实现
func watchChan(stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			fmt.Println("协程2接收到信号，stop!!!")
			return
		default:
			fmt.Println("2 watching.......")
			time.Sleep(2 * time.Second)
		}
	}
}

// 使用上下文context实现
func watchCtx(stop context.Context) {
	for {
		select {
		case <-stop.Done():
			fmt.Println("协程1接收到信号，stop!!!")
			return
		default:
			fmt.Println("1 watching.......")
			time.Sleep(2 * time.Second)
		}
	}
}
