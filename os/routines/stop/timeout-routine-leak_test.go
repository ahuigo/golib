package stop

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// https://geektutu.com/post/hpg-timeout-goroutine.html
func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		close(done)
		return fmt.Errorf("timeout")
	}
}

// timeout 导致routine泄漏: 1000个子协程没有退出, 且被done阻塞
/*
建议三选1：
1. done chan 加上缓冲区：不会阻塞routine
2. close(done)：routine执行 done <- true 也不会阻塞，不过会导致panic -- send on closed channel(read 不受影响)
3. doGoodthing(推荐): 通过select default来避免阻塞
*/
func TestTimeoutRoutineLeak(t *testing.T) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		timeout(doBadthing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine()) // 1002 // 1000个子协程没有退出
}

func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}
