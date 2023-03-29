package d

import (
	"fmt"
	"sync"
	"testing"
)

type LazyInt func() int

func Make(f func() int) LazyInt {
	var v int
	var once sync.Once //once只能用一次Do
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
