package cache2go_test

import (
	"testing"

	"github.com/13sai/go-learing/cache2go"
)


func TestDemo(t *testing.T) {
	cache2go.Demo()
}

func TestSetDataLoader(t *testing.T) {
	cache2go.SetDataLoader("sai0556")
}

func TestCallback(t *testing.T) {
	cache2go.Callback("t1", "sai")
}