package designPattern_test

import (
	"testing"

	"github.com/13sai/go-learing/designPattern"
)

func TestSubject_Notify(t *testing.T) {
	sub := &designPattern.Subject{}
	sub.Register(&designPattern.Obsever1{})
	sub.Register(&designPattern.Obsever2{})
	sub.Notify("hello 13sai")
}