//链接：https://www.zhihu.com/question/425625461/answer/2356045938

package demo

import "testing"

func barPanic() (r int) {
	defer func() {
		r += 4
		if recover() != nil {
			r += 8
		}
	}()

	var f func()
	defer f()    // undefined nil
	f = func() { // never used
		r += 2
	}

	return 1
}

func TestReturnPanic(t *testing.T) {
	println(barPanic())
}
