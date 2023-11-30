package gotest

import (
	"testing"
)

var _isHashKey map[any]int

func isHashableKey(key any) (canHash bool) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		canHash = false
	// 	}
	// }()
	_ = _isHashKey[key] // read: concurrent read ok
	// _isHashKey[key] = key.(int) // fatal: concurrent map write error
	return true
}

// go test -v -bench=. ./map-read-write_test.go
func Benchmark_Parralel_map_read(b *testing.B) {
	fn := func(goid, index int) bool {
		k := goid + index
		r := isHashableKey(k)
		return r
	}
	// RunParallel will create GOMAXPROCS goroutines
	// and distribute work among them.
	goid := 0
	b.RunParallel(func(pb *testing.PB) {
		goid++
		i := goid // GOMAXPROCS goroutines(num of cpu)
		j := 0
		for pb.Next() {
			j++
			fn(i, j)
		}
	})
}
