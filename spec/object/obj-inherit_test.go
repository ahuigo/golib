package demo

import (
	"fmt"
	"testing"
)

type AnimalBase struct {
	age int
}

func (p *AnimalBase) Play() {
	fmt.Println("showc age:", p.age)
}
func (p *AnimalBase) Info() {
	fmt.Printf("data: %#v\n", p)
}

// PetX 继承AnimalBase
type PetX struct {
	AnimalBase
	Name string
}

func TestInherit(t *testing.T) {
	pf := fmt.Printf
	petx := PetX{AnimalBase: AnimalBase{10}, Name: "DogAlex"}
	petx.AnimalBase.age = 1
	petx.age = 2 // 类型js prototype(实际修改的是 AnimalBase.age)

	pf("age: %#v,%#v\n", petx.AnimalBase.age, petx.age) // 2,2
	petx.Play()
	petx.Info() // 显示的是 AnimalBase 的数据, 不是PetX 的数据

}
