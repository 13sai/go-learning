package lock

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestSpinLock(t *testing.T) {
	var a uint32 = 1

	go func() {
		t.Log("1111", atomic.CompareAndSwapUint32(&a, 1, 2))
	}()

	go func() {
		t.Log("2222", atomic.CompareAndSwapUint32(&a, 1, 2))
	}()

	go func() {
		t.Log("3333")
		if !atomic.CompareAndSwapUint32(&a, 1, 2) {
			runtime.Gosched()
		}
		t.Log("3333222")
	}()

	t.Log("555", atomic.CompareAndSwapUint32(&a, 1, 2))

	time.Sleep(10 * time.Millisecond)
}
