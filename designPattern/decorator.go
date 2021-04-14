package designPattern

import "fmt"

type IDecorate interface {
	Do()
}

// 装饰器:实现接口，又定义了自己的事件DecorateFun，相当于抽象类
type Decorate struct {
	// 待装饰的抽象类
	Decorate IDecorate
}

func (d *Decorate) DecorateFun(i IDecorate) {
	d.Decorate = i
}

func (d *Decorate) Do() {
	if d.Decorate != nil {
		d.Decorate.Do()
	}
}

type DecorateA struct {
	Base Decorate
}

// 重写方法
func (d *DecorateA) Do() {
	fmt.Println("do DecorateA")
	d.Base.Do()
}
