package pkg

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
注意：
1. multi　goroutine受memory 限制:  一个goroutine大概占用2k内存，所以一般来说1G内存可以开启50万个goroutine
2. 系统资源limit: ulimit -a
	2.1　too many open files: ulimit -n:
*/
// 直接限制goroutine数量
func TestMultiGoroutineNum(t *testing.T) {
	var goroutineNum = 2
	lineChan := make(chan string, 10)
	// 1. 消费数据
	var wg sync.WaitGroup
	for i := 0; i < goroutineNum; i++ {
		wg.Add(1)
		go func() {
			// range chan
			defer wg.Done()
			for line := range lineChan {
				println(line)
				time.Sleep(time.Millisecond * 500)
			}
			println("goroutine closed")
		}()
	}

	// 2. 生产数据
	ticker := time.NewTicker(time.Millisecond * 10)
	// var ch1 <-chan int
	for i := 0; i < 3; i++ {
		<-ticker.C
		lineChan <- fmt.Sprintf("line %s", time.Now().String())
	}
	close(lineChan) // 读取range chan继续
	println("chan is closed")
	wg.Wait()
}

// 限制chan　缓冲区大小, 不限制goroutine数量
func TestMultiGoroutineChan(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		// 1. 生产数据()
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			t.Log(i)
			time.Sleep(time.Second)
			// 2. 消费数据
			<-ch
		}(i)
	}
	wg.Wait()
}
