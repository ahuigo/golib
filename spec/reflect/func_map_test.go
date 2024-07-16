package main

// refer: https://coolshell.cn/articles/21164.html
import (
	"fmt"
	"reflect"
	"testing"
)

func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

/*
*
额外展示了ref_func 的方法

	refn.Call([]reflect.Value{...})[0]
*/
func transform(slice, function interface{}, inPlace bool) interface{} {

	//check the `slice` type is Slice
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	//check the function signature
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	// MakeSlice
	sliceOutTypeVal := sliceInType
	if !inPlace {
		sliceOutTypeRef := reflect.SliceOf(fn.Type().Out(0)) // string ->  []string
		sliceOutTypeVal = reflect.MakeSlice(sliceOutTypeRef, sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		result := fn.Call([]reflect.Value{sliceInType.Index(i)})[0]
		sliceOutTypeVal.Index(i).Set(result)
	}
	return sliceOutTypeVal.Interface()

}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	//Check it is a funciton
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}
func TestFnCallMap(t *testing.T) {
	list := []string{"1", "2", "3", "4", "5", "6"}
	TransformInPlace(list, func(a string) string {
		return a + a + a
	})
	fmt.Printf("list:%#v\n", list)
	list2 := Transform(list, func(a string) string {
		return a + a + a
	})
	fmt.Printf("list2(%T):%#v\n", list2, list2)
}
