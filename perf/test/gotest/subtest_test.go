package gotest

import "testing"

func TestMul(t *testing.T) {
	// 用于组织测试用例
	t.Run("pos", func(t *testing.T) {
		if false {
			t.Fatal("fail")
		}

	})
	t.Run("neg", func(t *testing.T) {
		// t.Helper()
		if 6 != -6 {
			t.Fatal("fail")
		}
	})
}
