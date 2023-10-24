package demo

import (
	"fmt"
	"testing"
)

type None struct{}

// interface{} 可以跟其它类型比较
func TestInterfaceCompare(t *testing.T) {
	var a interface{} = "a"
	var b interface{} = "a"
	var c string = "a"
	m := map[interface{}]struct{}{}
	m[b] = None{}
	_, exist := m[a]
	fmt.Println(a == b)   // true
	fmt.Println(a == "a") // true
	fmt.Println(a == c)   // true
	fmt.Println(exist)    // true
}
