package gotest

import "testing"

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if false {
			t.Fatal("fail")
		}

	})
	t.Run("neg", func(t *testing.T) {
		if 6 != -6 {
			t.Fatal("fail")
		}
	})
}
