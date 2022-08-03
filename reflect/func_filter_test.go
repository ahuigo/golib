package main

import (
	"fmt"
	"reflect"
	"testing"
)

/**
注意：
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i])) //会把有值放到最前面,  有点像jsdom inplace替换： list.append(i)
	}
**/
func TestFilterInt(t *testing.T) {
	strs := []int{10, 0, 3, 0, 20}
	FilterInPlace(&strs, func(a int) bool {
		return a > 3
	})
	fmt.Println("out:", strs)
}

func Filter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

func FilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type, " +
			"not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	fmt.Println(in.Elem())
	in.Elem().SetLen(n)
	fmt.Println(in.Elem())
}

func filter(slice, function interface{}, inPlace bool) (interface{}, int) {

	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	boolType := reflect.ValueOf(true).Type()
	if !verifyFuncSignature(fn, elemType, boolType) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}

	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	out := sliceInType

	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i])) //排序
	}

	return out.Interface(), len(which)
}
