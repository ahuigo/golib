package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/timandy/routine"
)

func TestGoid(t *testing.T) {
	goid := routine.Goid()
	fmt.Printf("cur goid: %v\n", goid)
	go func() {
		goid := routine.Goid()
		fmt.Printf("sub goid: %v\n", goid)
	}()

	// 等待子协程执行完。
	time.Sleep(time.Second)
}

var threadLocal = routine.NewThreadLocal[string]()
var inheritableThreadLocal = routine.NewInheritableThreadLocal[string]()

// 以下代码简单演示了ThreadLocal的创建、设置、获取、跨协程传播等：
func TestShareThreadCache(t *testing.T) {
	threadLocal.Set("hello world")
	inheritableThreadLocal.Set("Hello world2")
	fmt.Println("threadLocal:", threadLocal.Get())
	fmt.Println("inheritableThreadLocal:", inheritableThreadLocal.Get())

	// 子协程无法读取之前赋值的“hello world”。
	go func() {
		fmt.Println("threadLocal in goroutine:", threadLocal.Get())
		fmt.Println("inheritableThreadLocal in goroutine:", inheritableThreadLocal.Get())
	}()

	// 但是，可以通过 Go/GoWait/GoWaitResult 函数启动一个新的子协程，当前协程的所有可继承变量都可以自动传递。
	routine.Go(func() {
		fmt.Println("threadLocal in goroutine by Go:", threadLocal.Get())
		fmt.Println("inheritableThreadLocal in goroutine by Go:", inheritableThreadLocal.Get())
	})

	// 等待子协程执行完。
	time.Sleep(time.Second)
}
