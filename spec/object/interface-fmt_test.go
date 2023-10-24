package demo

import (
	"fmt"
	"testing"
)

func TestInterfaceFmt(t *testing.T) {
	m := map[string]interface{}{"a": 1}
	// 自动识别int
	fmt.Printf("a=%T\n", m["a"].(int))
	fmt.Printf("a=%T\n", m["a"])
	fmt.Printf("m=%v,%T\n", m, m)

	b := map[string]interface{}{"b": []string{"str1"}}
	// 自动识别[]string
	fmt.Printf("b: type=%#T, value=%#v\n", b["b"], b["b"])

	var n1 interface{} = 1.25
	var n2 interface{} = float64(2222222222222222.22333333333)
	fmt.Printf("n1=%f,%T\n", n1, n1)
	fmt.Printf("n1=%v,%T\n", n1, n1)
	fmt.Printf("n2=%f,%T\n", n2, n2)
	fmt.Printf("n2=%v,%T\n", n2, n2)
}
