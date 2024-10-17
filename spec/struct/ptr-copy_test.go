package main

import (
	"fmt"
	"testing"
)

func TestPtrCopy(t *testing.T) {
	p := Person{name: "Alex"}
	var p2 *Person
	*p2 = p // nil 是不能赋值的
	p2.name = "Alex2"
	fmt.Println(p.name)
}
