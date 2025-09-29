package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// BenchmarkBufferedMultiGoroutine: 有缓冲 channel (大小 100), 多个 Goroutine 读写
func BenchmarkBufferedMultiGoroutine(b *testing.B) {
	// 100 个缓冲区:1282+w QPS(8worker), 3000w QPS(1worker)
	// 0个缓冲区:500+w QPS(8worker), 900w QPS(1worker)
	ch := make(chan int, 100)
	numWorkers := 8 // 模拟 8 个消费者
	count := int32(0)
	wg := sync.WaitGroup{}
	sum := int32(0)

	// 启动消费者
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range ch {
				atomic.AddInt32(&count, 1)
				atomic.AddInt32(&sum, int32(i))
			}
		}()
	}

	// 生产者
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i // Send
		}
		close(ch) // 完成后关闭 channel
	}()
	wg.Wait()
	if count != int32(b.N) {
		b.Errorf("Expected %d messages, but got %d", b.N, count)
	}
	println("Sum:", sum)
}
