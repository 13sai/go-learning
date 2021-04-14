package designPattern_test

import (
	"testing"

	"github.com/13sai/go-learing/designPattern"
)

func TestGenerate(t *testing.T) {
	factory := new(designPattern.Factory)
	p1 := factory.Generate("pro1")
	p1.This()

	p2 := factory.Generate("pro2")
	p2.This()
}

func TestAbstractFactory(t *testing.T) {
	f := new(designPattern.Factory1)
	p1 := f.CreatePro1()
	p1.This()

	p2 := f.CreatePro2()
	p2.This()
}
