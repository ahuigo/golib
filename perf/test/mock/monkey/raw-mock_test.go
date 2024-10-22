package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// 原始函数
func Echo(in string) string {
	return in
}

// 替换函数
func MockEcho(in string) string {
	return "mocked: " + in
}

// 替换函数实现
func replaceFunction(target, replacement interface{}) {
	targetValue := reflect.ValueOf(target).Elem()
	replacementValue := reflect.ValueOf(replacement).Elem()

	// 获取目标函数的指针
	targetPtr := unsafe.Pointer(targetValue.UnsafeAddr())
	replacementPtr := unsafe.Pointer(replacementValue.UnsafeAddr())

	// 替换函数指针
	*(*uintptr)(targetPtr) = *(*uintptr)(replacementPtr)
}

func TestRawMock(t *testing.T) {
	// 打印原始函数的结果
	fmt.Println("Original Echo:", Echo("hello"))

	// 替换函数实现
	replaceFunction(Echo, MockEcho)

	// 打印被 mock 后的函数结果
	fmt.Println("Mocked Echo:", Echo("hello"))
}
