package demo

import (
	"fmt"
	"testing"
)

// 泛型三元运算表达式
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func TestTernary(t *testing.T) {
	a := 5
	b := 10
	result := Ternary(a > b, a, b)
	fmt.Println("The larger value is:", result)

	str1 := "Hello"
	str2 := "World"
	longestString := Ternary(len(str1) > len(str2), str1, str2)
	fmt.Println("The longer string is:", longestString)
}
