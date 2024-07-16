package ptr

import (
	"fmt"
	"testing"
)

func TestSlicePtr(t *testing.T) {
	var a = []int{1, 2, 3}
	func(l *[]int) {
		*l = (*l)[1:]
	}(&a)
	fmt.Println(a) // 2,3
}

func TestSliceAppend(t *testing.T) {
	var scores []int
	fmt.Println(scores == nil, scores) // true,[]
	type Stu struct {
		scores []int
	}
	stu := Stu{scores: scores}
	func() {
		stu.scores = append(stu.scores, 1) // 不会改变外部的scores（不是指针）
	}()
	fmt.Println(stu.scores) // 1
	fmt.Println(scores)     // []
}

func TestSlicePtrAppend(t *testing.T) {
	var scores []int
	fmt.Println(scores == nil, scores) // true,[]
	type Stu struct {
		scores *[]int
	}
	stu := Stu{scores: &scores}
	func() {
		*stu.scores = append(*stu.scores, 1)
	}()
	fmt.Println(stu.scores) // &[1]
	fmt.Println(scores)     // &[1]
}
