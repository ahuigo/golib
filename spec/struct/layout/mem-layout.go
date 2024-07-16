package main

import (
	"fmt"
	"unsafe"
)

type Args struct {
	num1 int64
	num2 int8
}

type Flag struct {
	num1 int16
	num2 int32
}

type demo3 struct {
	c int32
	a struct{}
}

type demo4 struct {
	a struct{}
	c int32
}

func main() {
	fmt.Println(unsafe.Sizeof(demo3{})) // 8
	fmt.Println(unsafe.Sizeof(demo4{})) // 4
	fmt.Println(unsafe.Alignof(Args{}))
	fmt.Println(unsafe.Alignof(Flag{}))
}
