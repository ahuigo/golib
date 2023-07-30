package lock

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	for k := 0; k < 10; k++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
	println("done")
}
