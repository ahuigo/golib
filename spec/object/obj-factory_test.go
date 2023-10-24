/*
*
多态：相同的接口，大家都可以接入。工厂方法只是一种方式，不关心子类
参考：https://www.zybuluo.com/phper/note/1059734
作者：痕无落 链接：https://www.jianshu.com/p/b333c5f34ef6
*/
package demo

import (
	"fmt"
	"testing"
)

type Animal interface {
	Age() int
	Type() string
}

// Cat 实现接口Animal
type Cat struct {
	MaxAge int
}

func (o *Cat) Age() int {
	return o.MaxAge
}
func (o *Cat) Type() string {
	return "Cat"
}

// Dog 实现接口Animal
type Dog struct {
	MaxAge int
}

func (o *Dog) Age() int {
	return o.MaxAge
}
func (o *Dog) Type() string {
	return "Dog"
}

// 用接口Animal 或者 interface{} 实现多态Type()/Age()
func Factory(name string) Animal {
	switch name {
	case "dog":
		return &Dog{MaxAge: 20}
	case "cat":
		return &Cat{MaxAge: 10}
	default:
		panic("No such animal")
	}
}

func TestObjFactory(t *testing.T) {
	animal := Factory("dog")
	fmt.Printf("%s max age is: %d", animal.Type(), animal.Age())
	// Dog max age is: 20
}
