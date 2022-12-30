package demo

import (
	"fmt"
	"testing"
)

// go run -gcflags=-G=3 %
func print[T any](arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}
func TestAny(t *testing.T) {
	strs := []string{"Hello", "World", "Generics"}
	decs := []float64{3.14, 1.14, 1.618, 2.718}
	nums := []int{2, 4, 6, 8}
	print(strs)
	print(decs)
	print(nums)
}

// any 是接口类型约束
// 还有：comparable
