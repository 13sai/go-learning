package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChan2(t *testing.T) {
	stop := make(chan struct{})

	done := make(chan struct{}, 2)

	for i := 0; i < 2; i++ {
		go run(stop, done)
	}

	time.Sleep(2 * time.Second)

	close(stop)

	for i := 0; i < 2; i++ {
		<-done
	}
}

func run(stop <-chan struct{}, done chan<- struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("stop...")
			done <- struct{}{}
			return
		case <-time.After(time.Second):
			fmt.Println("gfhdjhg")
		}
	}
}
