package perf

import (
	"fmt"
	"testing"
)

/**
使用空struct　优化内存
- 集合Set: map 实现value 为 struct{} 时，可以节省内存
- 不发送数据的信道(channel)
*/

func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something")
	close(ch)
}

func TestEmptyChan(t *testing.T) {
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
}
