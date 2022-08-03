package main

import (
    "fmt"
)

type Pet struct {
    age int
}
func (p *Pet) Play() {
    fmt.Println("showc age:", p.age)
}

type PetX struct {
    Pet
    name string
}
func main() {
    pf := fmt.Printf

    petx:=PetX{Pet:Pet{20}}
    petx.Pet.age=1
    petx.age=2

    pf("%#v,%#v\n", petx.Pet.age, petx.age)
    petx.Play()

}
