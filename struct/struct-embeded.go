package main
 
import (
	"fmt"
)
 
type Person struct {
	Name        string `label:"Person Name: " uppercase:"true"`
	Age         int    `label:"Age is: "`
	Sex         string `label:"Sex is: "`
	Description string
}
type P struct {
    Person
	Age         int    `label:"Age is: "`
}
 
func main() {
	person := P{
        Person:Person{
            Name:        "Tom",
            Age:         29,
            Sex:         "Male",
            Description: "Cool",
        },
	}
	fmt.Printf("%#v\n", person)
    fmt.Printf("age:%#v", person.Age) //0
}
