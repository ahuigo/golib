package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/**
配合演示了ref_slice的如下方法：
	refv.Len()
	refv.Index(i)
	reflect.SliceOf(type)
	reflect.MakeSlice(sliceOutType, vdata.Len(), vdata.Len())
	item.Set(refv)

额外展示了ref_func 的方法
	refn.Call([]reflect.Value{...})[0]
*/
func TestSliceLoop(t *testing.T) {
	Map := func(data interface{}, fn interface{}) interface{} {
		vfn := reflect.ValueOf(fn)
		vdata := reflect.ValueOf(data)
		// vdata.Len() vdata.Index(i) 是refv slice 具有的

		// make slice
		sliceOutType := reflect.SliceOf(vdata.Type().Elem()) // string ->  []string
		sliceOutTypeVal := reflect.MakeSlice(sliceOutType, vdata.Len(), vdata.Len())

		// loop slice
		for i := 0; i < vdata.Len(); i++ {
			item := vdata.Index(i)
			rfParams := []reflect.Value{item}
			result := vfn.Call(rfParams)[0]
			sliceOutTypeVal.Index(i).Set(result)
			// replace inplace
			item.Set(result)
		}
		return sliceOutTypeVal.Interface()
	}

	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := Map(strs, strings.ToUpper)
	fmt.Printf("strs:%#v\n", strs)
	fmt.Printf("upstrs:%#v\n", upstrs)
}
