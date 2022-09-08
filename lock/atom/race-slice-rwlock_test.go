package main

import (
	"sync"
	"testing"
	"time"
)

// 1. 读写之间互斥(读时可能会阻塞)，读与读不互拆
func TestSliceRwLock(t *testing.T) {
	type Obj struct {
		count int
	}

	s := []*Obj{
		{1},
		{2},
	}
	lock := sync.RWMutex{}
	// writer
	go func() {
		i := 1
		for {
			i += 1
			lock.Lock()
			s = []*Obj{
				{1},
				{2},
			}
			time.Sleep(100 * time.Nanosecond)
			lock.Unlock()
		}
	}()

	var wg sync.WaitGroup
	// read
	i := 1
	since := time.Now()
	for {
		i += 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.RLock()
			s1 := s
			if s1[1].count != 2 {
				panic("error")
			}
			time.Sleep(100 * time.Nanosecond)
			lock.RUnlock()
		}()
		if i > 1e6 {
			break
		}
	}
	wg.Wait()
	diff := time.Since(since)
	t.Log(diff)
}
