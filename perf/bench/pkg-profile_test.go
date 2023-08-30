package gotest

import (
	"math/rand"
	"testing"

	"github.com/pkg/profile"
)

func TestPkgProfileMem(t *testing.T) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	randomString := func(n int) string {
		b := make([]byte, n)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		return string(b)
	}

	plusConcat := func(n int, str string) string {
		s := ""
		for i := 0; i < n; i++ {
			s += str
		}
		return s
	}
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	var str = randomString(10)
	// profile 分析不知道这个函数名，实现显示为：main.func1
	plusConcat(10000, str)
}
