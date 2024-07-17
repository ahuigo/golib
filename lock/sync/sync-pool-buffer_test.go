package lock

/**
A Pool is a set of temporary objects that may be individually saved and retrieved.

Any item stored in the Pool may be removed automatically at any time without notification. If the Pool holds the only reference when this happens, the item might be deallocated.

A Pool is safe for use by multiple goroutines simultaneously.
*/

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

func TestSyncPollBufferNoRelease(t *testing.T) {
	var bufPool = sync.Pool{
		New: func() interface{} {
			// The Pool's New function should generally only return pointer
			return new(bytes.Buffer)
		},
	}
	acquireBuffer := func() *bytes.Buffer {
		return bufPool.Get().(*bytes.Buffer)
	}

	consumeBuffer := func() {
		// bufPool 获取了一个 *bytes.Buffer，但是你没有将其放回 bufPool。这意味着 bufPool 不再持有对该 *bytes.Buffer 的引用
		buf := acquireBuffer()
		buf.WriteString("hello")
		fmt.Println(buf.Len())
		// 一旦 consumeBuffer 函数执行完毕，buf 变量就会离开其作用域，因此对 *bytes.Buffer 的引用就会被丢弃。 gc就会回收它
	}
	for i := 0; i < 100; i++ {
		consumeBuffer()
	}
}

// get reset put
func TestSyncPollBufferBadRelease(t *testing.T) {
	var bufPool = sync.Pool{
		New: func() interface{} {
			// The Pool's New function should generally only return pointer
			return new(bytes.Buffer)
		},
	}
	acquireBuffer := func() *bytes.Buffer {
		return bufPool.Get().(*bytes.Buffer)
	}
	releaseBuffer := func(buf *bytes.Buffer) {
		if buf != nil {
			buf.Reset()
			bufPool.Put(buf)
		}
	}
	// 1.
	buf := acquireBuffer()
	buf.WriteString("ahui")
	a := buf.Bytes()
	releaseBuffer(buf) // 错误：buf 变量在这里被放回pool，但是 a 变量仍然持有对 *bytes.Buffer 的引用

	// 2.
	buf2 := acquireBuffer()
	buf2.WriteString("hello2")
	b := buf2.Bytes()
	releaseBuffer(buf)

	// test
	fmt.Println("a:", string(a))
	fmt.Println("b:", string(b))

	/*
		sync.Pool 两种做法：
		1.　要么只get，不reset+put放回。等引用消失后gc会自动回收
		1.　要么只get，使用完后, 再reset+put放回
		3.　不能是get，再reset+put放回，再继续使用
	*/

}
