package fxdemo

import (
	"testing"

	"go.uber.org/dig"
)

// 如果注册的方法返回的参数是可以为nil的，可以使用option来实现
func TestProvideOption1(t *testing.T) {
	type Student struct {
		dig.Out
		Name string
		Age  *int `option:"false"` // option:"false" 表示可以为nil
	}

	c := dig.New()
	if err := c.Provide(func() Student {
		return Student{
			Name: "ahuigo",
		}
	}); err != nil {
		t.Fatal(err)
	}

	if err := c.Invoke(func(n string, age *int) {
		t.Logf("name: %s", n)
		if age == nil {
			t.Log("age is nil")
		} else {
			t.Logf("age: %d", age)
		}
	}); err != nil {
		t.Fatal(err)
	}
}

/*
输出
name: ahuigo
age is nil
*/
