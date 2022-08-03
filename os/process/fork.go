package main

import (
	"log"
	"os"
    "time"

	"github.com/docker/docker/pkg/reexec"
)

func init() {
	log.Printf("init start, os.Args = %+v\n", os.Args)
	reexec.Register("childProcess", childProcess)
	if reexec.Init() {
		os.Exit(0)
	}
}

func childProcess() {
	log.Println("childProcess")
}

func main() {
	log.Printf("main start, os.Args = %+v\n", os.Args)
    //  reexec.Command 可以修改子进程的 os.Args[0]，所以子进程会直接找到 reexec.Init 上reexec.Register 所注册的函数，然后执行，返回true，最后调用 os.Exit(0)
	cmd := reexec.Command("childProcess")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Panicf("failed to run command: %s", err)
	}
	if err := cmd.Wait(); err != nil {
		log.Panicf("failed to wait command: %s", err)
	}
    time.Sleep(100000000000)
	log.Println("main exit")
}
