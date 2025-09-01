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
	println(1)
	_ = c
    println("1.2")

}

func TestSetFinalizer(t *testing.T) {
    println(0)
	fooCache()
    println("2")
	cs := make([]int, 1e7)
    println("3")
	runtime.GC() // call stopJanitor
    println("4")
	println(len(cs))
}
