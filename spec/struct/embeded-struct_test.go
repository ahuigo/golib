package main

import (
	"fmt"
	"testing"
)

func TestEmbedStruct(t *testing.T) {
	type PersonA struct {
		Name        string `label:"Person Name: " uppercase:"true"`
		Age         int    `label:"Age is: "`
		Sex         string `label:"Sex is: "`
		Description string
	}
	type P struct {
		PersonA
		Age int `label:"Age is: "`
	}
	person := P{
		PersonA: PersonA{
			Name:        "Tom",
			Age:         29,
			Sex:         "Male",
			Description: "Cool",
		},
	}
	fmt.Printf("%#v\n", person)
	fmt.Printf("age:%#v", person.Age) //0
}
