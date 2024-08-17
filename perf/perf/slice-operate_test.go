package perf

import (
	"fmt"
	"testing"
)

// 参考：post/go/go-array.md
func TestSliceCopy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3)
	// method 1
	copy(b, a) // 不报错
	fmt.Println(b)

	// method 2
	b = append([]int(nil), a...)
	fmt.Println(b)
}

func TestSliceDelete(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	// 删除第3个元素
	i := 2
	// a = append(a[:i], a[i+1:]...)
	a = a[:i+copy(a[i:], a[i+1:])]
	fmt.Println(a)
}

func TestSliceDeleteGC(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	// 删除第3个元素
	i := 2
	copy(a[i:], a[i+1:])
	a[len(a)-1] = 0  // or zero value of T
	a = a[:len(a)-1] //触发gc回收
	fmt.Println(a)
}
