package main

import (
	"fmt"
	"sync"
)

func Smap() {
	var sm sync.Map

	sm.Store("www", "2343125")
	sm.Store("test", 111)

	var v interface{}

	v, _ = sm.Load("www")
	fmt.Printf("%T", v)
	sm.Delete("test")
	v, _ = sm.Load("test")
	fmt.Printf("%T", v)
}
