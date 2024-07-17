package pkg

import (
	"sync"
	"testing"
	"time"
)

func multi_run(count int, fn func()) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
	println(count, "tasks done")

}

func sleep() {
	n := uint(1)
	time.Sleep(time.Duration(n) * time.Second)
	println(n, "s elasped!")

}

// 可以用for　range　chan 实现给并行任务,发task
func TestMultiRun(t *testing.T) {
	fn := sleep
	multi_run(5, fn)
}
