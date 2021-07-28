package nsq

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func TestReceive(t *testing.T) {
	topic := "sai"
	go receive(topic)
	send(topic)
	// send(topic)

	shut := make(chan os.Signal)
	signal.Notify(shut, os.Interrupt)
	fmt.Println("开始监听信号")
	<-shut
}
