package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd, err := exec.LookPath("ls")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}

	var args []string
	args =  append(args, "aaa")
	sh := exec.Command(cmd, []string{"-a"}...)
	// Start 开始执行但不返回
	if err := sh.Start(); err!=nil {
		fmt.Println("err -- ", err)
	}
	err = sh.Wait()
	fmt.Println("err -- ", err)

	fmt.Println(1<<16)
}