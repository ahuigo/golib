package main

import (
	"sync/atomic"
)

func ExampleCompareAndSwapInt64() {
	var i *int64 = new(int64)
	*i = 3
	atomic.CompareAndSwapInt64(i, 3, 6)
	println(*i)
}
