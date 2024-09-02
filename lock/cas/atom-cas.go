// https://github.com/golang/go/issues/17604
// 无锁队列的实现 https://coolshell.cn/articles/8239.html
// @source go-lib/goroutine/cas/atom-cas.go
package queue

import (
	"sync"
	"sync/atomic"
)

const sizeSection int = 12000 * 10000

type Message struct {
	id int
}

type Queue struct {
	mu   *sync.Mutex
	db   []Message
	hasp int32
	tail int
	size int
}

func (q *Queue) Len() int {
	return len(q.db)
}

func New() Queue {
	q := Queue{
		new(sync.Mutex), make([]Message, sizeSection, sizeSection), 0, 0, sizeSection,
	}
	return q
}

// 这里应该优化成*Queue，否则copy
func (q Queue) PushTailCAS(n Message) {
	for {
		if atomic.CompareAndSwapInt32(&q.hasp, 0, 1) {
			break
		}
	}
	q.db[q.tail] = n
	q.tail++
	if q.tail > q.size {
		q.db = append(q.db, make([]Message, sizeSection, sizeSection)...)
		q.size += sizeSection
	}
	q.hasp = 0
	//atomic.StoreInt32(&q.hasp, 0)
}
func (q Queue) PushTailMutex(n Message) {
	q.mu.Lock()
	q.db[q.tail] = n
	q.tail++
	if q.tail > q.size {
		q.db = append(q.db, make([]Message, sizeSection, sizeSection)...)
		q.size += sizeSection
	}
	q.mu.Unlock()
}

// fixed
/*
Qestion: CompareAndSwapInt32 为什么比mutex更慢呢？因为q　逃逸到堆栈上
 1. 堆栈调用的开销:
 	在堆栈上的调用atomic.CompareAndSwapInt32 更慢(mutex虽然会调用，但是不会分配一个新的对象并创建写屏障（write barrier）)
    CompareAndSwapInt32　在compiler中，会 allocates a new object and creates write barrier
 2. The compiler moves q to the heap because atomic.CompareAndSwapInt32 is implemented in assembly.
    Escape analysis has to conservatively assume that any pointer passed to assembly escapes.
    转义分析必须保守地假设传递给程序集的任何指针都会转义。即它们可能被函数外部引用q。这可能导致对象被分配在堆上，而不是在栈上
*/
func (q *Queue) PushTailCASFixed(n Message) {
	for {
		if atomic.CompareAndSwapInt32(&q.hasp, 0, 1) {
			break
		}
	}
	q.db[q.tail] = n
	q.tail++
	if q.tail == q.size {
		q.db = append(q.db, make([]Message, sizeSection, sizeSection)...)
		q.size += sizeSection
	}
	q.hasp = 0
}

func (q *Queue) PushTailMutexFixed(n Message) {
	q.mu.Lock()
	q.db[q.tail] = n
	q.tail++
	if q.tail == q.size {
		q.db = append(q.db, make([]Message, sizeSection, sizeSection)...)
		q.size += sizeSection
	}
	q.mu.Unlock()
}
