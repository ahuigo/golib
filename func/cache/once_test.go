package d

import (
	"fmt"
	"sync"
	"testing"
)

type LazyInt func() int

/*
sync.Once 作用:
1.　缓存
2.　并发阻塞: refer to /lock/sync/once
*/
func Make(f func() int) LazyInt {
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
	n := Make(func() int { return 23 }) // Or something more expensive…
	fmt.Println(n())                    // Calculates the 23
	fmt.Println(n() + 42)               // Reuses the calculated value
}
