package queue

import (
	"fmt"
	"testing"
	_ "time"
)

type MessageX struct {
	id int
}

func (q MessageX) Set() {
	q.id = 123
	q2 := &q
	q2.Print()
}
func (q *MessageX) Print() {
	fmt.Println(q.id)
}
func TestMessage(t *testing.T) {
	m := MessageX{id: 10}
	m.Set()
	m.Print()

}

/*
Fixed: 优化成*Queue，否则会有copy
Qestion: CAS 为什么更慢呢？
 1. 在堆栈上的调用atomic.CompareAndSwapInt32，就错了
    CompareAndSwapInt32 会 compiler allocates a new object and creates write barrier
 2. The compiler moves q to the heap because atomic.CompareAndSwapInt32 is implemented in assembly.
    Escape analysis has to conservatively assume that any pointer passed to assembly escapes.
    转义分析必须保守地假设传递给程序集的任何指针都会转义。

BenchmarkPushTailCAS-10           	45875161	        26.31 ns/op	      64 B/op	       1 allocs/op
BenchmarkPushTailMutex-10         	85902454	        13.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkPushTailCASFixed-10      	90363997	        20.10 ns/op	      56 B/op	       0 allocs/op
BenchmarkPushTailMutexFixed-10    	63611265	        16.45 ns/op	      52 B/op	       0 allocs/op
*/
func BenchmarkPushTailCAS(b *testing.B) {
	fmt.Println("size0")
	b.StopTimer()
	q := New()
	b.StartTimer()
	m := 0
	for i := 0; i < b.N; i++ {
		m = i
		q.PushTailCAS(Message{id: i})
	}
	fmt.Println("size1:", q.Len(), m)
}

func BenchmarkPushTailMutex(b *testing.B) {
	b.StopTimer()
	q := New()
	b.StartTimer()
	m := 0
	for i := 0; i < b.N; i++ {
		m = i
		q.PushTailMutex(Message{id: i})
	}
	fmt.Println("size2:", q.Len(), m)
}

func BenchmarkPushTailCASFixed(b *testing.B) {
	b.StopTimer()
	q := New()
	b.StartTimer()
	m := 0
	for i := 0; i < b.N; i++ {
		m = i
		q.PushTailCASFixed(Message{id: i})
	}
	fmt.Println("size3:", q.Len(), m)
}

func BenchmarkPushTailMutexFixed(b *testing.B) {
	b.StopTimer()
	q := New()
	b.StartTimer()
	m := 0
	for i := 0; i < b.N; i++ {
		m = i
		q.PushTailMutexFixed(Message{id: i})
	}
	fmt.Println("size4:", q.Len(), m)
}
