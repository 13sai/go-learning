package designPattern_test

import (
	"testing"

	"github.com/13sai/go-learing/designPattern"
)

func TestChain(t *testing.T) {
	adH := &designPattern.AdHandle{}
	senH := &designPattern.SensitiveHandle{}
	adH.H = senH

	adH.Handle("hello 13sai")
}
