package multi

import (
	"fmt"
	"testing"
)

type MultiType1 interface {
	func(int) int | ~func() int |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
type MultiType2[K1 any, V any] interface {
	func() V | ~func(K1) V | ~func(int, int) int |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func FnMultiType1[T MultiType1](a T) T {
	return a
}
func FnMultiType2[K1 any, V any, T MultiType2[K1, V]](a T) T {
	return a
}
func FnMultiType3[
	K1 any,
	V any,
	T int | func(K1) V,
](a T) T {
	return a
}

func TestMultiType(t *testing.T) {
	a := FnMultiType1(func(a int) int {
		return a
	})
	fmt.Println(a(1))
	_ = FnMultiType2[int, bool](func(a int) bool {
		return a > 0
	})
	_ = FnMultiType3[int, bool](func(a int) bool {
		return a > 0
	})
}
