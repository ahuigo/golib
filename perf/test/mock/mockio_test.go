package m

import (
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
)

type myInterface interface {
	Foo(a int) int
}

func TestSimple(t *testing.T) {
	SetUp(t)
	m := Mock[myInterface]()
	WhenSingle(m.Foo(Any[int]())).ThenReturn(42)
	_ = m.Foo(10)
	Verify(m, AtLeastOnce()).Foo(10)
}
