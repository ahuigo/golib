package fxdemo

import (
	"testing"

	"go.uber.org/dig"
)

// refer to : https://www.cnblogs.com/li-peng/p/14708132.html
/**
如果有很多相同类型的返回参数，可以把他们放在同一个slice里，和命名方式一样，有两种使用方式
第一种在调用Provide时直接使用dig.Group
*/
func TestDigGroupInvokeIn(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}
	NewUser := func(name string, age int) func() *Student {
		return func() *Student {
			return &Student{name, age}
		}
	}
	container := dig.New()
	if err := container.Provide(NewUser("tom", 3), dig.Group("stu")); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUser("jerry", 1), dig.Group("stu")); err != nil {
		t.Fatal(err)
	}
	//invoke dig.In with group
	type inParams struct {
		dig.In
		StudentList []*Student `group:"stu"` //dig.In 告诉dig StudentList的值由group:stu 提供
	}
	Info := func(params inParams) error {
		if len(params.StudentList) == 0 {
			t.Fatalf("StudentList is empty")
		}
		for _, u := range params.StudentList {
			t.Log(u.Name, u.Age)
		}
		return nil
	}
	if err := container.Invoke(Info); err != nil {
		t.Fatal(err)
	}
}

// 或者不用dig.Group/In, 而是使用结构体嵌入dig.Out来实现，tag里要加上了group标签
func TestDigGroupProvideOut(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}
	type Rep struct {
		dig.Out
		//这个flatten的意思是，底层把组表示成[]*Student，如果不加flatten会表示成[][]*Student
		StudentList []*Student `group:"stu,flatten"`
	}
	NewUser := func(name string, age int) func() Rep {
		return func() Rep {
			r := Rep{}
			r.StudentList = append(r.StudentList, &Student{
				Name: name,
				Age:  age,
			})
			return r
		}
	}

	container := dig.New()
	if err := container.Provide(NewUser("tom", 3)); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUser("jerry", 1)); err != nil {
		t.Fatal(err)
	}
	type InParams struct {
		dig.In

		StudentList []*Student `group:"stu"`
	}
	Info := func(params InParams) error {
		for _, u := range params.StudentList {
			t.Log(u.Name, u.Age)
		}
		return nil
	}
	if err := container.Invoke(Info); err != nil {
		t.Fatal(err)
	}
}
