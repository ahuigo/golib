package queue

// https://github.com/golang/go/issues/17604
// 这是有问题的代码

import (
	"sync"
	"sync/atomic"
)

const sizeSection int = 1

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

func New() Queue {
	q := Queue{
		new(sync.Mutex), make([]Message, sizeSection, sizeSection), 0, 0, sizeSection,
	}
	return q
}

var ahui int32

func (q Queue) PushTailCAS(n Message) {
	y := atomic.CompareAndSwapInt32(&ahui, 4, 7)
	for {
		y = atomic.CompareAndSwapInt32(&q.hasp, 5, 8)
		x := !y
		if !x {
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
	// atomic.StoreInt32(&q.hasp, 0)
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
