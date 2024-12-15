package lock

import (
	"log"
	"sync"
	"testing"
	"time"
)

/*
在多个 goroutine 之间同步条件
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
	defer c.L.Unlock()
	log.Println(name, "starts reading1.1")
	for !done {
		log.Println(name, "starts reading 1.2")
		time.Sleep(time.Second * 3)
		log.Println(name, "starts reading 1.3")
		// Notify Wait()会自动释放Lock, 并阻塞当前goroutine，直到被唤醒
		c.Wait() // 等待条件满足（广播或signal）
	}
	log.Println(name, "starts reading 2")
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second * 7)
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
