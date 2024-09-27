package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func mergeStruct1[T any](s1 *T) *T {
	d1, _ := json.Marshal(s1)
	// 利用反射根据s1的类型创建一个新的*T(s1 类型*T, Elem()获取类型T)
	s := reflect.New(reflect.TypeOf(s1).Elem()).Interface()
	json.Unmarshal(d1, s)
	return s.(*T)
}

// rtype.Elem(): 获取ptr 实际的类型
func TestMergeStruct(t *testing.T) {
	type Opt struct {
		Name     string
		Pass     string
		PoolSize int
	}
	s1 := Opt{Name: "name1", Pass: "pass1"}
	s := mergeStruct1(&s1)
	fmt.Printf("%#v\n", s)
	t.Log(s)
}
