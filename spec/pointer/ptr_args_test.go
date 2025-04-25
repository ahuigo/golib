package demo

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestPtrArgs(t *testing.T) {
	fn := func(s *string) {
		t.Log(s == nil) //false
		*s = "hello"
	}
	var s string
	fn(&s)
	t.Log(s)
}

func TestArgAssignVal(t *testing.T) {
	fn := func(a *int, b *int) {
		*a = *b
	}
	a := 1
	b := 2
	fn(&a, &b)
	b = 3
	if a == b {
		t.Error("expect a != b ")
	}
	t.Log(a, b)

}

func TestArgAssignPtr(t *testing.T) {
	fn := func(a **int, b *int) {
		// 将 b 的地址赋值给 *a, (改变的不是a指针，而是a 指向的变量指针*a)
		*a = b
	}
	var a *int
	b := 2
	c := 3

	// method 1
	fn(&a, &b)
	b = 4
	if *a != b {
		t.Error("expect a == b")
	}
	t.Log(*a, b, c)

	// method 2
	SetAnyPointerUnsafe(&a, &c)
	c = 5
	if *a != c {
		t.Error("expect a == c")
	}
	t.Log(*a, b, c)

}

func SetPointer2[T any](a **T, b *T) {
	*a = b
}

func SetAnyPointer(a, b any) {
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)

	// 检查 a 是否是指针的指针，b 是否是指针
	if aValue.Kind() != reflect.Ptr || aValue.Elem().Kind() != reflect.Ptr {
		panic("a must be a pointer to a pointer")
	}
	if bValue.Kind() != reflect.Ptr {
		panic("b must be a pointer")
	}

	// 将 b 赋值给 *a
	aValue.Elem().Set(bValue)
}

func SetAnyPointerUnsafe(a, b any) {
	ap := (*uintptr)(unsafe.Pointer(
		reflect.ValueOf(a).Pointer(),
	))
	bp := reflect.ValueOf(b).Pointer()
	*ap = bp
}
