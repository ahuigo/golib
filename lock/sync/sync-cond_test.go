package lock

import (
	"log"
	"sync"
	"testing"
	"time"
)

/*
sync.Cond 基于互斥锁/读写锁，它和互斥锁的区别是什么呢？
场景：一个生产者生产数据，多个goroutine消费者等待同一个条件(广播)，当条件满足时，一起消费数据
*/

func TestSyncCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)
	time.Sleep(time.Second * 3)
}

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait() // 等待条件满足（广播或signal）
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")

	wakeAll := true
	if wakeAll {
		// 唤醒所有等待的goroutine
		c.Broadcast() // 通知所有等待的goroutine　消费
	} else {
		// 只唤醒一个的goroutine
		c.Signal()
	}

}
