package main

import (
	"fmt"
	"testing"
)

func TestArgCopyStructAny(t *testing.T) {
    type Stu struct{
        age int
    }
    s1:=Stu{1}
	f := func(a any) {
        s1.age=100
        fmt.Printf("copy struct any:%v, old:1\n", a.(Stu).age)
	}
	f(s1)
}
func TestArgPassSliceAny(t *testing.T) {
	ids := []int{1, 2}
	f := func(a any) {
        ids[0]=100
        fmt.Printf("pass slice any:%v\n", a.([]int)[0])
	}
	f(ids)
}

func TestArgPassSlice(t *testing.T) {
	f := func(ids []int) {
        ids[0]=100
	}
	ids := []int{1}
	f(ids)
	fmt.Printf("pass slice ids:%v, len(ids)=%d\n", ids, len(ids))
}

func TestArgCopyArrayAny(t *testing.T) {
	ids := [2]int{1, 2}
	f := func(a any) {
        ids[0]=100
        fmt.Printf("copy array any:%v\n", a.([2]int)[0])
	}
	f(ids)
}
func TestArgCopyArray(t *testing.T) {
	f := func(ids [2]int) {
        ids[0]=100
	}
	ids := [2]int{1, 2}
	f(ids)
	fmt.Printf("array ids:%v, len(ids)=%d\n", ids, len(ids))

}

