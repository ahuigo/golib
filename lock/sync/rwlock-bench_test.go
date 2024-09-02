package lock

import (
	"sync"
	"testing"
	"time"
)

/*
*
conlusion:
 1. 读多：读写锁更快
 2. 写多：读写锁稍快
 3. 读写相等：读写锁稍快

BenchmarkReadMore-10       	     381	   3078471 ns/op	  113851 B/op	    2018 allocs/op
BenchmarkReadMoreRW-10     	    2074	    563497 ns/op	  112453 B/op	    2005 allocs/op
BenchmarkWriteMore-10      	     384	   3120714 ns/op	  113361 B/op	    2015 allocs/op
BenchmarkWriteMoreRW-10    	     416	   2857492 ns/op	  113494 B/op	    2015 allocs/op
BenchmarkEqual-10          	     388	   3081700 ns/op	  113458 B/op	    2016 allocs/op
BenchmarkEqualRW-10        	     692	   1732109 ns/op	  112830 B/op	    2009 allocs/op
*/
func BenchmarkReadMore(b *testing.B)    { benchmark(b, &Lock{}, 9, 1) }
func BenchmarkReadMoreRW(b *testing.B)  { benchmark(b, &RWLock{}, 9, 1) }
func BenchmarkWriteMore(b *testing.B)   { benchmark(b, &Lock{}, 1, 9) }
func BenchmarkWriteMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 1, 9) }
func BenchmarkEqual(b *testing.B)       { benchmark(b, &Lock{}, 5, 5) }
func BenchmarkEqualRW(b *testing.B)     { benchmark(b, &RWLock{}, 5, 5) }

type RW interface {
	Write()
	Read()
}
type Lock struct {
	count int
	mu    sync.Mutex
}

const cost = time.Microsecond

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *RWLock) Read() {
	l.mu.RLock()
	_ = l.count
	time.Sleep(cost)
	l.mu.RUnlock()
}

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
