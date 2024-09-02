package perf

import (
	"testing"
)

/*
Refer: https://geektutu.com/post/hpg-for-range.html
优化建议：
1. for index 的性能大约是 range (同时遍历下标和值) 的 许多 倍:  因为　range　会复制一份数组，而 for 只是引用
1. 如果struct 换成*struct，for/range 性能差不多（不会发生复制）
*/
type Item struct {
	id  int
	val [4096]byte
}

func BenchmarkForStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeIndexStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// for index 的性能大约是 range (同时遍历下标和值) 的 许多 倍
// 因为　range　会复制一份数组，而 for 只是引用
func BenchmarkRangeStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
