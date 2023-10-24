package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetValue(t *testing.T) {
	var i float64 = 3.1
	// 不传指针就不能改变值 reflect.ValueOf(i).SetFloat(7.4)
	reflect.ValueOf(&i).Elem().SetFloat(7.4)
	reflect.Indirect(reflect.ValueOf(&i)).SetFloat(7.4)
	fmt.Printf("%v\n", i)
}
