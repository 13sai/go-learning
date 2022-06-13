package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const (
	cgroupMemoryGierarchyMount = "/sys/fs/cgroup/memory"
	exe                        = "/proc/self/exe"
)

func main() {
	if os.Args[0] == exe {
		fmt.Printf("current pid %d\n", syscall.Getgid())

		cmd := exec.Command("sh", "-c", "stress --vm-bytes 100m --vm-keep -m 1")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}

	cmd := exec.Command(exe)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	fmt.Println(cmd.Process.Pid)

	// 创建子进程
	newCgroup := path.Join(cgroupMemoryGierarchyMount, "cg-memory")
	if err := os.Mkdir(newCgroup, 0755); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(path.Join(newCgroup, "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(path.Join(newCgroup, "memory.limit_in_bytes"), []byte("100m"), 0644); err != nil {
		panic(err)
	}

	cmd.Process.Wait()
}
