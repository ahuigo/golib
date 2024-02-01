package test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestKill(t *testing.T) {
	// 创建一个接收信号的通道
	sigs := make(chan os.Signal, 1)

	// 使用 signal.Notify 注册要接收的信号: 所有信号
	signal.Notify(sigs)
	// 或指定信号
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGCHLD)

	// 使用 goroutine 来异步处理信号
	go func() {
		for {
			sig := <-sigs
			// convert sig to number
			fmt.Println("Received signal:", sig, int(sig.(syscall.Signal)))
		}
	}()

	fmt.Println("Program is running. Press Ctrl+C to exit.", os.Getpid())
	// 使主进程不会退出
	select {}
}
