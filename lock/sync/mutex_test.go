package lock

import (
	"sync"
	"testing"
	"time"
)

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/*
Mutex: https://colobu.com/2018/12/18/dive-into-sync-mutex/
互斥锁有两种状态：正常状态和饥饿状态。

	在正常状态下，所有等待锁的 goroutine 按照FIFO顺序等待。唤醒的 goroutine 不会直接拥有锁，而是会和新请求锁的 goroutine 竞争锁的拥有。新请求锁的 goroutine 具有优势：它正在 CPU 上执行，而且可能有好几个，所以刚刚唤醒的 goroutine 有很大可能在锁竞争中失败。在这种情况下，这个被唤醒的 goroutine 会加入到等待队列的前面。 如果一个等待的 goroutine 超过 1ms 没有获取锁，那么它将会把锁转变为饥饿模式。

	在饥饿模式下，锁的所有权将从 unlock 的 goroutine 直接交给交给等待队列中的第一个。新来的 goroutine 将不会尝试去获得锁，即使锁看起来是 unlock 状态, 也不会去尝试自旋操作，而是放在等待队列的尾部。

	如果一个等待的 goroutine 获取了锁，并且满足一以下其中的任何一个条件：(1)它是队列中的最后一个(表明已经没有G要等待/阻塞了)；(2)它等待的时候小于1ms(表明之前的G能很快获取锁)。它会将锁的状态转换为正常状态。

	正常状态有很好的性能表现，饥饿模式也是非常重要的，因为它能阻止尾部延迟的现象(表明大量的G在等待锁)

关于读写锁 vs 互斥锁的性能：

	读写比为 9:1 时，读写锁的性能约为互斥锁的 3 倍
	读写比为 1:9 时，读写锁性能相当
	读写比为 5:5 时，读写锁的性能约为互斥锁的 2 倍
*/
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
