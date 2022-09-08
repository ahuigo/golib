package lock

import (
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var mutex = &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func(j int) {
			mutex.Lock()
			defer mutex.Unlock()
			println(j)
			time.Sleep(time.Second * 1)
		}(i)
	}
	time.Sleep(time.Second * 5)
}
