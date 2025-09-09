package perf

import (
	"runtime"
	"testing"
	"unsafe"
)

//go:noinline
func toStringUnsafe(b []byte) string {
	// 避免内存拷贝
	return unsafe.String(unsafe.SliceData(b), len(b))
}

//go:noinline
func toStringNormal(b []byte) string {
	return string(b)
}

func BenchmarkUnsafeString(b *testing.B) {
	b.Run("unsafe", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bs := []byte("hello, unsafe")
			_ = toStringUnsafe(bs)
		}
	})
	b.Run("normal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bs := []byte("hello, unsafe")
			a := toStringNormal(bs)
			_ = a + " world"
		}
	})
}

/*
*
unsafe.Pointer 是 GC 可知的：如果一个 unsafe.Pointer 变量本身是可达的（例如，它在栈上或在另一个可达的堆对象中），那么它所指向的内存对象也被认为是可达的，GC 不会回收它。unsafe.Pointer 就像一个通用类型的“安全指针”，GC 会跟踪它。

uintptr 是 GC 不可知的：uintptr 是一个整数类型，足以存放一个指针的地址值。当你将一个 unsafe.Pointer 转换为 uintptr 后，GC 就无法再识别它是一个指针了。它只把 uintptr 变量看作一个普通的整数。如果此时没有其他“安全指针”或 unsafe.Pointer 指向原始的内存对象，那么该对象就可能被 GC 回收
*/

func TestUnitPtrToStruct(t *testing.T) {
	type MyStruct struct {
		A int
		B string
		C float64
	}

	s := &MyStruct{A: 10, B: "hello", C: 3.14}
	ptr := uintptr(unsafe.Pointer(s))

	// 强制GC，演示 uintptr 不被GC跟踪，可能导致对象被回收
	runtime.GC()

	// 尝试转换回来并访问
	s2 := (*MyStruct)(unsafe.Pointer(ptr)) // s2 会被gc跟踪

	// 注意：如果对象被GC回收，这里可能panic或行为未定义
	// 在正常情况下，由于s还在作用域，GC不会回收，但这是演示
	if s2.A != s.A || s2.B != s.B || s2.C != s.C {
		t.Errorf("Expected %+v, got %+v", s, s2)
	}
}
func TestUnsafePointerToStruct(t *testing.T) {
	type MyStruct struct {
		A int
		B string
		C float64
	}

	s := &MyStruct{A: 10, B: "hello", C: 3.14}
	ptr := unsafe.Pointer(s)

	s2 := (*MyStruct)(ptr)

	if s2.A != s.A || s2.B != s.B || s2.C != s.C {
		t.Errorf("Expected %+v, got %+v", s, s2)
	}
}
