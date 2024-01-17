package main

import (
	"math/rand"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func TestPkgProf(t *testing.T) {
	// 1. mem.pprof
	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// 2. cpu.pprof
	//defer profile.Start().Stop()
	// testing 支持生成 CPU、memory 和 block 的 profile 文件。
	// 3. 也可以使用 testing 包的 -cpuprofile 和 -memprofile 标志来生成 profile 文件。
	// -cpuprofile=$FILE -memprofile=$FILE, -memprofilerate=N 调整记录速率为原来的 1/N。
	concat(100)
}
