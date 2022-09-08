package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 交替打印字符
// go-lib/goroutine/string-atom-race.go
// 其实应该用读写锁改进:
// 1. 读写之间互斥(读时可能会阻塞)，读与读不互拆
func TestRaceStringLock(t *testing.T) {
	const (
		FIRST  = "WHAT THE"
		SECOND = "F*CK"
	)
	var mutex = &sync.Mutex{}
	var s string
	go func() {
		i := 1
		for {
			i = 1 - i
			mutex.Lock()
			if i == 0 {
				s = FIRST
			} else {
				s = SECOND
			}
			mutex.Unlock()
			time.Sleep(10 * time.Nanosecond)
		}
	}()

	for {
		mutex.Lock()
		fmt.Println(s)
		mutex.Unlock()
		time.Sleep(10 * time.Nanosecond)
	}
}
