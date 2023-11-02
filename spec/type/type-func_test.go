package main

import (
	"fmt"
	"testing"
)

type Option interface {
	apply(string)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(string) //它是一个类型，可接受并初始化为结构体函数

func (f optionFunc) apply(log string) {
	f(log) //调用结构体函数
}

func log(o Option) {
	fmt.Printf("option arg=%v\n", o)
}

func TestTypeFunc(t *testing.T) {
	// //它是一个类型，可接受并初始化为函数
	option := optionFunc(func(name string) {
		fmt.Printf("option.name=%v\n", name)
	})

	fmt.Println("-----测试 optionFunc------")
	option.apply("ahui")
	option("ahui2")
	log(option)

	fmt.Println("-----测试 Option------")
	var op2 Option = option
	op2.apply("ahui(op2)")
	// op2("ahui(op2)") 不可以
}
