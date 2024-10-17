package demo

import (
	"fmt"
	"testing"
)

func TestConvertStruct(t *testing.T) {
	type Cat1 struct {
		age  int
		Name string
	}
	type Cat2 struct {
		age  int
		Name string
	}
	// copy struct mehtod:
	//1. sp:=s
	//2. sp:=Type(s)
	cat1 := Cat1{age: 7}
	cat2 := Cat2(cat1)
	cat2.age = 1

	fmt.Println(cat1)
	fmt.Println(cat2)
}
