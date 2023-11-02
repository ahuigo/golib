package main

import (
	"fmt"
	"testing"
)

type StringArray []string

func TestTypeArray(t *testing.T) {
	var i StringArray = []string{"a"}
	fmt.Printf("%T\n", i)

	func(j []string) {
		fmt.Printf("%T\n", j)
	}(i)
}
