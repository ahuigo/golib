package main

import (
	. "sync/atomic"
	"testing"
)

func TestValue(t *testing.T) {
	var v Value
	if v.Load() != nil {
		t.Fatal("initial Value is not nil")
	}
	v.Store(42)
	x := v.Load()
	if xx, ok := x.(int); !ok || xx != 42 {
		t.Fatalf("wrong value: got %+v, want 42", x)
	}
	v.Store(84)
	x = v.Load()
	if xx, ok := x.(int); !ok || xx != 84 {
		t.Fatalf("wrong value: got %+v, want 84", x)
	}
}
