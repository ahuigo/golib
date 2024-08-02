package main

import (
	"sync/atomic"
	"testing"
)

func TestCount1(t *testing.T) {
	var counter atomic.Int32
	counter.Add(1)
	counter.Add(3)
	counter.Add(-1)
	t.Log(counter.Load())
}

func TestCount2(t *testing.T) {
	var counter uint32
	atomic.AddUint32(&counter, 1)
	t.Log(counter)
}
