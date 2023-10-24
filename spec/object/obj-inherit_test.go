package demo

import (
	"fmt"
	"testing"
)

type PetInheritBase struct {
	age int
}

func (p *PetInheritBase) Play() {
	fmt.Println("showc age:", p.age)
}

// PetX 继承PetInheritBase
type PetX struct {
	PetInheritBase
	name string
}

func TestInherit(t *testing.T) {

	pf := fmt.Printf

	petx := PetX{PetInheritBase{20}, "Alex"}
	petx.PetInheritBase.age = 1
	petx.age = 2 // 类型js prototype

	pf("%#v,%#v\n", petx.PetInheritBase.age, petx.age) // 2,2
	petx.Play()

}
