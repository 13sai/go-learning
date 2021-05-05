package designPattern

import "fmt"

var aa = "a"

// 简单工厂
type Factory struct{}

type Product interface {
	This()
}

type Pro1 struct{}

func (p Pro1) This() {
	fmt.Println("This is pro1")
}

type Pro2 struct{}

func (p Pro2) This() {
	fmt.Println("This is pro2")
}

func (f Factory) Generate(name string) Product {
	switch name {
	case "pro1":
		return Pro1{}
	case "pro2":
		return Pro2{}
	default:
		return nil
	}
}

type ProInter interface {
	This()
}

// 抽象工厂
type AbstractFactory interface {
	CreatePro1() ProInter
	CreatePro2() ProInter
}

type Factory1 struct{}

func (f *Factory1) CreatePro1() ProInter {
	return new(Pro1)
}

func (f *Factory1) CreatePro2() ProInter {
	return new(Pro2)
}
