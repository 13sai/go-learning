package designPattern

import "sync"

type Instance struct {
	Str string
}

var singletonA, singletonB *Instance

func GetInstanceA() *Instance {
	if singletonA == nil {
		singletonA = &Instance{"A"}
	}
	return singletonA
}

var once sync.Once

func GetInstanceB() *Instance {
	once.Do(func() {
		singletonB = &Instance{"B"}
	})
	return singletonB
}
