package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
sync.Once 作用:
1.　缓存
2.　并发阻塞: refer to /lock/sync/once
*/
var out = 0

func TestSyncOnce(t *testing.T) {
	var once sync.Once
	onceBodyFunc := func(i int) func() {
		return func() {
			fmt.Println("once start:", i)
			time.Sleep(5 * time.Second)
			fmt.Println("Only once")
			out += i * 100
		}
	}
	onceBody1 := onceBodyFunc(1)
	onceBody2 := onceBodyFunc(2)
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println("routine step1:", j)
			once.Do(onceBody1) // 阻塞等待(	o.m.Lock())
			fmt.Println("routine step 2:", j)
			once.Do(onceBody2) // 已经done，不会再执行了
			fmt.Println("end:", j, ", out:", out)
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

/*
sync.Once 作用:
1.　缓存
2.　并发阻塞: refer to /lock/sync/once
*/
type LazyInt func() int

func MakeCacheFunc(f func() int) LazyInt {
	var v int
	var once sync.Once //once只能用一次Do, 同时调用do 被阻塞
	return func() int {
		println("call func")
		once.Do(func() {
			v = f()
			f = nil // so that f can now be GC'ed
		})
		return v
	}
}

func TestOnceMake(t *testing.T) {
	n := MakeCacheFunc(func() int {
		return 23
	}) // Or something more expensive…
	fmt.Println(n())      // Calculates the 23
	fmt.Println(n() + 42) // Reuses the calculated value
}
