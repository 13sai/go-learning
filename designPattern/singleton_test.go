package designPattern_test

import (
	"testing"

	"github.com/13sai/go-learing/designPattern"
)

func TestGetInstanceA(t *testing.T) {
	t.Log("GetInstanceA", designPattern.GetInstanceA())
}

func TestGetInstanceB(t *testing.T) {
	t.Log("GetInstanceB", designPattern.GetInstanceB())
	t.Log("GetInstanceB", designPattern.GetInstanceB())
}
