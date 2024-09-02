package perf

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// https://geektutu.com/post/hpg-string-concat.html
/**

 */
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

/*
cap: 16 + 32 + 64 + ... + 122880 = 0.52 MB
1. strings.Builder 性能比 bytes.Buffer 略快约 10% 。strings.Builder 直接将底层的 []byte 转换成了字符串类型直接返回了回来。
BenchmarkBuilderConcat-10          36043             32989 ns/op          106496 B/op          1 allocs/op
BenchmarkPreByteConcat-10          37802             30504 ns/op          212992 B/op          2 allocs/op
BenchmarkBufferConcat-10           30471             41626 ns/op          212992 B/op          2 allocs/op
*/
func TestBuilderConcatCap(t *testing.T) {
	var str = randomString(10)
	var builder strings.Builder
	cap := 0
	for i := 0; i < 10000; i++ {
		if builder.Cap() != cap {
			fmt.Print(builder.Cap(), " ")
			cap = builder.Cap()
		}
		builder.WriteString(str)
	}
}

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}
func builderConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str)) // 提前grow
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	buf.Grow(n * len(s)) // 提前grow
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}
func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

// go test -bench="Concat$" -benchmem .
func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, builderConcat) }
func BenchmarkPreByteConcat(b *testing.B) { benchmark(b, preByteConcat) }
func BenchmarkBufferConcat(b *testing.B)  { benchmark(b, bufferConcat) }
