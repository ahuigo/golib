package main

import "fmt"
type Person struct {
    name string
    age int
}

type People map[string]Person

func main() {
    p := make(People)
    p["HM"] = Person{"Hank McNamara", 39}
    //p["HM"].age = p["HM"].age + 1 //cannot assign
    fmt.Printf("age: %d\n", p["HM"].age)
}
