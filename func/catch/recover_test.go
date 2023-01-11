package main

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func g(i int) {
	fmt.Println("Panic within g()!")
	panic(i)
}

func f() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover rtn:%v, type:%T\n", r, r)
			fmt.Println("stacktrace from panic: \n  " + string(debug.Stack()))
			err = fmt.Errorf("recover exception: %v", r)
		}
	}()

	i := 10
	fmt.Println("Calling g with ", i)
	g(i)
	fmt.Println("Returned normally from g.")
	return nil
}
func TestM(t *testing.T) {
	err := f()
	fmt.Printf("main rtn: %v\n", err)

}
