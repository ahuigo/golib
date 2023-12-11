package main

import (
	"fmt"
	"testing"
)

func foo() (int, int) {
	return 1, 2
}

func TestArgLocal(t *testing.T) {
	res, err := func() (res int, err int) {
		res1, err := foo()
		if err > 0 {
			// err 返回局部变量2!
			return
		}
		return res1 * 100, err

	}()
	fmt.Println(res, err)

}
