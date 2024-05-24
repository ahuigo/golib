package perf

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

var ssTable sync.Map

func stringptr(s string) uintptr {
	// unsafe.String or unsafe.StringData
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

func StringIntern(s string) string {
	actual, _ := ssTable.LoadOrStore(s, s)
	return actual.(string)
}

func TestStrIntern(t *testing.T) {
	f := func() string {
		s0 := "hello world1"
		s1 := fmt.Sprintf("hello world%d", 1)
		fmt.Printf("%v\n", stringptr(s0) == stringptr(s1))
		// false

		// 复用同一个字符串，返回的是同一个地址
		fmt.Printf("%v\n", stringptr(StringIntern(s0)) == stringptr(StringIntern(s1)))
		// true

		// gc回收s0,s1
		s := StringIntern(s0)
		return s
	}
	StringIntern("hello world1")
	f()
}
