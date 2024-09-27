package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 反射对象 与 interface 对象转化
func TestVarConvert(t *testing.T) {
	pf := fmt.Printf
	var i float64 = 3.1
	var a interface{} = i

	// data -> reflect object
	v := reflect.ValueOf(i)

	// reflect object -> origin interface
	b := v.Interface()
	pf("a=b? %v\n", a == b) //true

}
