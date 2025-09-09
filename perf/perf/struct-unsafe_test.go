package perf

import (
	"testing"
	"unsafe"
)

type TestStruct struct {
	ID   int
	Name string
	Data [10240000]int // 增加大小以显示拷贝开销
}

//go:noinline
func getStruct(m map[string]*TestStruct, key string) *TestStruct {
	return m[key]
}

const bucketSize = 16

func BenchmarkDirectReturn(b *testing.B) {
	m := make(map[string]*TestStruct, bucketSize)
	for i := 0; i < bucketSize; i++ {
		key := string(rune(i))
		m[key] = &TestStruct{ID: i, Name: "test", Data: [10240000]int{i}}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := string(rune(i % bucketSize))
		v := getStruct(m, key) // 从函数返回大结构体
		_ = *v                 // 解引用以模拟使用
	}
}

// unsafe.Pointer 与 直接返回pointer 性能差一点点
func BenchmarkUnsafeReturn(b *testing.B) {
	m := make(map[string]*TestStruct, bucketSize)
	for i := 0; i < bucketSize; i++ {
		key := string(rune(i))
		m[key] = &TestStruct{ID: i, Name: "test", Data: [10240000]int{i}}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := string(rune(i % bucketSize))
		v := m[key] // pointer
		ptr := (*TestStruct)(unsafe.Pointer(v))
		_ = *ptr // Dereference to simulate usage
	}
}
