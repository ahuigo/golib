package main

import (
	"fmt"
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

/*
*
mock 派系分为两类：
1. monkey patch：基于修改内存态方法区汇编码，将需 mock 的方法修改为 jmp 指令，从而实现对全局方法的实现的替换。
2. codegen：使用工具生成 interface 的 mock 实现，其 mock 实现将接口方法代理到基于反射的一套 mock api 上。
bytedance/monkey 是 monkey patch 派系，而 goconvey 是 codegen 派系。
1. monkey是在运行时重写了函数指令
*/
func Foo(in string) string {
	return in
}

type A struct{}

func (a A) Foo(in string) string { return in }

var Bar = 0

// go test -gcflags="all=-l -N" -timeout 1800s -run '^TestMockFunc$' m2 -count=1 -v
func TestMockFunc(t *testing.T) {
	PatchConvey("TestMockXXX", t, func() {
		Mock(Foo).Return("c").Build()   // mock function
		Mock(A.Foo).Return("c").Build() // mock method
		MockValue(&Bar).To(1)           // mock variable

		So(Foo("a"), ShouldEqual, "c")        // assert `Foo` is mocked
		So(new(A).Foo("b"), ShouldEqual, "c") // assert `A.Foo` is mocked
		So(Bar, ShouldEqual, 1)               // assert `Bar` is mocked
	})
	// mock is released automatically outside `PatchConvey`
	fmt.Println(Foo("a"))        // a
	fmt.Println(new(A).Foo("b")) // b
	fmt.Println(Bar)             // 0
}
