package stop

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
1. unclose chan 导致routine泄漏: 结束时，子协程还在阻塞，多了一个
2. close(chan) 不会影响read：一个通道被其发送数据协程队列和接收数据协程队列中的所有协程引用着。因此，如果一个通道的这两个队列只要有一个不为空，则此通道肯定不会被垃圾回收。
*/
func TestRoutineLeak(t *testing.T) {
	t.Helper()
	t.Log(runtime.NumGoroutine())
	taskCh := make(chan int, 10)

	go do(taskCh)
	for i := 0; i < 100; i++ {
		taskCh <- i
	}
	// close(taskCh) // 记住！！！
	time.Sleep(time.Second)
	t.Log(runtime.NumGoroutine())
}

func do(taskCh chan int) {
	for {
		select {
		// case t := <-taskCh: // close　后，不会报错，t取到的就是0
		case t, ok := <-taskCh:
			if !ok {
				fmt.Println("done")
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}
