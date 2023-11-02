package main

import (
	"fmt"
	"reflect"
	"testing"
)

type MyInt int

func test1(i int) {
	fmt.Printf("%T\n", i)
}
func test2(i MyInt) {
	fmt.Printf("%T\n", i)
}

func TestTypeArg(t *testing.T) {
	var i int = 1
	var r rune
	var ri int32
	//int32,int32,main.MyInt
	fmt.Printf("is=?:%T,%T,%T\n", ri, r, MyInt(i))

	test1(i)             //int
	test1(int(MyInt(i))) //int
	test2(MyInt(i))      //error: test1(MyInt(i))

	var vi interface{} = ""
	fmt.Printf("interface type:%T\n", vi)
	fmt.Printf("reflect interface:%v\n", reflect.TypeOf(vi))
	fmt.Printf("reflect interface:%v\n", reflect.ValueOf(vi).Kind())

}
