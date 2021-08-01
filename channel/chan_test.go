package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch := make(chan bool, 0)
	go runGo(ch)
	go runGo(ch)

	ch <- true

	time.Sleep(2 * time.Second)
	close(ch)

}

func runGo(ch chan bool) {
	ok := <-ch
	fmt.Println("over", ok)
}
