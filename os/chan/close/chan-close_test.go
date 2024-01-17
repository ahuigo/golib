package demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChanClose(t *testing.T) {
	lineChan := make(chan string, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// line, ok := <-lineChan
		// fmt.Printf("goroutine closed:line=%#v, ok=%#v\n", line, ok)
		line := <-lineChan // 不会因为close报错
		fmt.Printf("goroutine closed:line=%#v, \n", line)
	}()
	time.Sleep(time.Millisecond * 2000)
	close(lineChan)
	wg.Wait()
	println("main is closed")
}


func TestChanRange(t *testing.T) {
	lineChan := make(chan string, 10)
    // 1. 消费数据
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// range chan
		defer wg.Done()
		for line := range lineChan {
			println(line)
			time.Sleep(time.Millisecond * 1000)
		}
		println("goroutine closed")
	}()

    // 2. 生产数据
	ticker := time.NewTicker(time.Millisecond * 10)
	// var ch1 <-chan int
	for i := 0; i < 3; i++ {
		<-ticker.C
		lineChan <- fmt.Sprintf("line %s", time.Now().String())
	}
	close(lineChan) // 读取range chan继续
	println("chan is closed")
	wg.Wait()

}
