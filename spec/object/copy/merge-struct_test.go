package demo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func mergeStruct1[T any](s1, s2 *T) *T {
	d1, _ := json.Marshal(s1)
	d2, _ := json.Marshal(s2)
	// 利用反射根据s1的类型创建一个新的*T(s1 类型*T, Elem()获取类型T)
	s := reflect.New(reflect.TypeOf(s1).Elem()).Interface()
	json.Unmarshal(d2, s)
	json.Unmarshal(d1, s) // 空格值会被覆盖
	return s.(*T)
}
func mergeStruct2[T any](s1, s2 T) T {
	return s1
}
func TestMergeStruct(t *testing.T) {
	type Opt struct {
		Name     string
		Pass     string
		PoolSize int
	}
	s1 := Opt{Name: "name1", Pass: "pass1"}
	s2 := Opt{PoolSize: 2}
	d2, _ := json.Marshal(s2)
	json.Unmarshal(d2, &s1)

	s := mergeStruct1(&s1, &s2)
	fmt.Printf("%#v\n", s)
	t.Log(s)
}
