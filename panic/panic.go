package main

import (
	"fmt"
	"time"
)

func main() {
	pChan := make(chan interface{})

	fmt.Println("---")
	go func() {
		defer func() {
			if p := recover(); p != nil {
				pChan <- p
			}
		}()
	}()
	panic("sjfgdsj")
	time.Sleep(1 * time.Second)
	fmt.Println(<-pChan)
}
