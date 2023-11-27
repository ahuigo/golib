package lock

import (
	"sync"
	"testing"
	"time"
)

func TestRWMutex(t *testing.T) {
	var mutex = &sync.RWMutex{}
	for i := 0; i < 5; i++ {
		go func(j int) {
			println(getTime(), "start:", j)
			mutex.Lock() // 阻塞的
			println(getTime(), "start2:", j)
			defer mutex.Unlock()
			println(j)
			time.Sleep(time.Second * 1)
		}(i)
	}
	time.Sleep(time.Second * 5)
}

func TestRWMutexTry(t *testing.T) {
	var mutex = &sync.RWMutex{}
	for i := 0; i < 5; i++ {
		go func(j int) {
			println(getTime(), "start:", j)
			isLock := mutex.TryLock() // 非阻塞的
			if !isLock {
				println(getTime(), "start1:", j)
				mutex.Lock()
			}
			println(getTime(), "start2:", j)
			defer mutex.Unlock()
			println(j)
			time.Sleep(time.Second * 1)
		}(i)
	}
	time.Sleep(time.Second * 5)
}
