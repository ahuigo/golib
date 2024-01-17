package main

import (
	"reflect"
	"testing"
)

// 射创建对象的耗时约为 原生new 的 1.5 倍，相差不是特别大。
func BenchmarkReflectNew(b *testing.B) {
	type Config struct {
	}
	var config *Config
	typ := reflect.TypeOf(Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config, _ = reflect.New(typ).Interface().(*Config)
	}
	_ = config
}
