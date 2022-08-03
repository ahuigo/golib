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
 
 
func main() {
	p1 := Person{
		Name:        "Tom",
		Age:         29,
		Sex:         "Male",
		Description: "Cool",
	}
	p2 := Person{
		Name:        "Tom",
		Age:         29,
		Sex:         "Male",
		Description: "Cool",
	}
    fmt.Printf("%v", p1==p2)
 
}
