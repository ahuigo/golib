package lock

import (
	"sync"
	"testing"
	"time"
)

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func TestMutex(t *testing.T) {
	var mutex = &sync.Mutex{}
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
