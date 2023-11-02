package function

import (
	"testing"
)

type AnyOfThree[T any] struct {
	Value T
}

// T 是*AnyOfThree[any] 中的any
// T ~int 表示int 或者 基于int的类型(type MyInt int)
func genericArgs[T ~string | ~int | *AnyOfThree[any]](val T) {
	switch v := any(val).(type) {
	case string:
		println("string:", v)
	case int:
		println("int", v)
	default:
		println("default", v)
	}
}

type genericType interface {
	~string | ~int | *AnyOfThree[any]
}

func genericArgs2[T genericType](val T) {

}
func TestArgs(t *testing.T) {
	type MyInt int
	genericArgs(1)
	genericArgs(MyInt(1))
	genericArgs("str")
	m := &AnyOfThree[int]{
		Value: 1,
	}
	genericArgs(m.Value)
	genericArgs((*m).Value)
	genericArgs2((*m).Value)

}
