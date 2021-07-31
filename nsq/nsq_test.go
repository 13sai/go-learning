package nsq

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestReceive(t *testing.T) {
	topic := "sai0556"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go Send(ctx, cancel, topic)
	go Receive(ctx, cancel, topic)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	t.Log("开始监听")

	select {
	case <-ctx.Done():
		t.Log("ctx done")
		return
	case <-sig:
		t.Log("signal exit...")
		cancel()
		time.Sleep(2 * time.Second)
		return
	}

	// send(topic)

}
