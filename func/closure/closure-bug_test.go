package closure

import (
	"fmt"
	"testing"
	"time"
)

type Api struct {
	Topic string
}

func addMonitorAPI(api *Api) {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("trigger", api.Topic)
		fmt.Printf("api addr2:%p\n", api)
	}()
}

func TestClosureGoroutine(t *testing.T) {
	apis := []Api{
		{"topic1"},
		{"topic2"},
	}
	for i, api := range apis {
		_ = i
		_ = api
		// 方法1：api2 存的地址, 是api的地址; 存在闭包bug(循环结束后，api就是最后一个元素的副本)
		api2 := &api
		// 方法2： api2 存的指针, 指向slice元素; 没有闭包bug(循环结束后，指向slice元素)
		// api2 = &apis[i]
		fmt.Printf("api addr:%p\n", api2)
		addMonitorAPI(api2)
	}
	time.Sleep(2 * time.Second)
}

/*
*
VariableLoop

	3 3 3

ValueLoop

	0 1 2

VariableRange

	2 2 2

ValueRange

	0 1 2
*/
func TestClosureFor(t *testing.T) {
	VariableLoop := func() {
		f := make([]func(), 3)
		for i := 0; i < 3; i++ {
			// closure over variable i
			f[i] = func() {
				fmt.Println(i)
			}
		}
		fmt.Println("VariableLoop")
		for _, f := range f {
			f()
		}
	}

	ValueLoop := func() {
		f := make([]func(), 3)
		for i := 0; i < 3; i++ {
			i := i
			// closure over value of i
			f[i] = func() {
				fmt.Println(i)
			}
		}
		fmt.Println("ValueLoop")
		for _, f := range f {
			f()
		}
	}

	VariableRange := func() {
		f := make([]func(), 3)
		for i := range f {
			// closure over variable i
			f[i] = func() {
				fmt.Println(i)
			}
		}
		fmt.Println("VariableRange")
		for _, f := range f {
			f()
		}
	}

	ValueRange := func() {
		f := make([]func(), 3)
		for i := range f {
			i := i
			// closure over value of i
			f[i] = func() {
				fmt.Println(i)
			}
		}
		fmt.Println("ValueRange")
		for _, f := range f {
			f()
		}
	}

	VariableLoop()
	ValueLoop()
	VariableRange()
	ValueRange()
}
