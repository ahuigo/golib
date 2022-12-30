package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	assert.Equal(t, Identity(1), 1)
	assert.Equal(t, Identity("abc"), "abc")
}

func TestConst(t *testing.T) {
	assert.Equal(t, Const[int, string](1)("a"), 1)
	assert.Equal(t, Const[string, int]("a")(1), "a")
}

func TestFirstAndSecond(t *testing.T) {
	assert.Equal(t, First(1, "a"), 1)
	assert.Equal(t, First("a", 1), "a")
	assert.Equal(t, Second(1, "a"), "a")
	assert.Equal(t, Second("a", 1), 1)
}

func TestComposeAndThen(t *testing.T) {
	var (
		double = func(x int) int { return x + x }
		incr   = func(x int) int { return x + 1 }
	)

	assert.Equal(t, Compose(incr, double)(4), 9)
	assert.Equal(t, Compose(double, incr)(4), 10)
	assert.Equal(t, Then(incr, double)(4), 10)
	assert.Equal(t, Then(double, incr)(4), 9)
}

func TestCurry(t *testing.T) {
	var (
		add = func(x int, y int) int { return x + y }
	)
	assert.Equal(t, Curry2(add)(1)(2), 3)
	assert.Equal(t, UnCurry2(Curry2(add))(1, 2), 3)
}

func TestFlip(t *testing.T) {
	var (
		concat = func(a, b string) string { return a + b }
	)
	assert.Equal(t, Flip(concat)("abc", "def"), "defabc")
	assert.Equal(t, Flip(Flip(concat))("abc", "def"), "abcdef")
}

func TestFix(t *testing.T) {
	var (
		f = func(x int) int {
			if x%2 == 0 {
				return x / 2
			} else {
				return (x + 1) / 2
			}
		}
	)
	assert.Equal(t, Fix(f)(10), 1)
	assert.Equal(t, Fix(f)(0), 0)
}

func TestOn(t *testing.T) {
	var (
		g    = func(x, y string) int { return len(x) + len(y) }
		argF = func(x int) string { return fmt.Sprintf("%d", x) }
	)
	assert.Equal(t, On(g, argF)(10, 123), 5)
}

func TestCall(t *testing.T) {
	var (
		incr = func(a int) int { return a + 1 }
	)
	assert.Equal(t, Call(incr, 10), 11)
}
