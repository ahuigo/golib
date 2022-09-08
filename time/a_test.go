package main

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	t0 := time.Now()
	time.Sleep(time.Second * 2)
	diff := time.Since(t0)
	t.Log(diff, time.Second)
	if diff > 1*time.Second {
		t.Fatal("xxx")
	}
	t.Log(diff, time.Second)

}
