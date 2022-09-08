package main

import (
	"runtime"
	"testing"
)

func fooCache() {
	type Cache struct {
		count int
	}

	stopJanitor := func(c *Cache) {
		println("gc will call finalized")
		c.count = 0
	}
	c := &Cache{1}
	runtime.SetFinalizer(c, stopJanitor)
	println(0)
	_ = c

}

func TestSetFinalizer(t *testing.T) {
	fooCache()
	cs := make([]int, 1e7)
	runtime.GC()
	println(len(cs))
}
