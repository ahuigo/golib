package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestTypeAs(t *testing.T) {
	var target *os.PathError
	err := &os.PathError{
		Op: "open",
	}

	// errors.As(err, &target)
	As := func(err, targetPtr any) {
		tPtrVal := reflect.ValueOf(targetPtr)
		targetType := tPtrVal.Type().Elem()
		// 判断err是否赋值给targetType类型
		if reflect.TypeOf(err).AssignableTo(targetType) {
			tPtrVal.Elem().Set(reflect.ValueOf(err))
		}
	}
	As(err, &target)
	if target.Op != "open" {
		t.Errorf("target.Op = %q, want %q", target.Op, "open")
	}
}

type Foo struct {
	A int    `tag1:"First Tag" tag2:"Second Tag"`
	B string `tag1:"Tag1"`
}
type SliceStr []string

func TestExamType(t *testing.T) {
	sl := []int{1, 2, 3}
	greeting := "hello"
	greetingPtr := &greeting
	f := Foo{A: 10, B: "Salutations"}
	structPtr := &f
	sliceStr := SliceStr{}

	sliceType := reflect.TypeOf(sl)
	sliceStrType := reflect.TypeOf(sliceStr)
	strType := reflect.TypeOf(greeting)
	strPtrType := reflect.TypeOf(greetingPtr)
	structType := reflect.TypeOf(f)
	structPtrType := reflect.TypeOf(structPtr)

	examiner(strType, 0)
	examiner(strPtrType, 0)
	examiner(sliceType, 0)
	examiner(sliceStrType, 0)

	mapType := reflect.TypeOf(map[string]int{"age": 1, "height": 1})
	//interfaceType := reflect.TypeOf(map[string]interface{}{"age": 1, "name": "ahui"})
	examiner(mapType, 0)

	examiner(structType, 0)
	examiner(structPtrType, 0)
}

/*
type.Name: string, "", StructName,..
type.Kind: string, slice, struct,..
type.Elem(): 获取ptr 实际的类型
*/
func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth)+"TypeName:", t.Name(), ",Kind:", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1) + "---Elem type----")
		examiner(t.Elem(), depth+1)
	case reflect.Map:
		fmt.Println(strings.Repeat("\t", depth+1) + "---Key and Elem---")
		examiner(t.Key(), depth+1)
		examiner(t.Elem(), depth+1)
	case reflect.Interface:
		//todo
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Printf("%s field:%d==================%%T=%T\n", strings.Repeat("\t", depth+1), i+1, f)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "Name:", f.Name, ",Type:", f.Type.Name(), ",Kind:", f.Type.Kind())
			fmt.Printf("%s fieldByName:%%T=%T\n", strings.Repeat("\t", depth+1), t.FieldByName)
			fA, _ := t.FieldByName("A")
			fmt.Printf("%s fieldByName(A):%%T=%T\n", strings.Repeat("\t", depth+1), fA)
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "f.Tag:", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1:", f.Tag.Get("tag1"), ",tag2:", f.Tag.Get("tag2"))
			}
		}
	}
}
