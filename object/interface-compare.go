package main

import (
	"fmt"
)
type None struct{}


func main() {
    var a interface{} = "a"
    var b interface{} = "a"
    m := map[interface{}]struct{}{}
    m[b] =None{}
    _,exist := m[a]
    fmt.Println(a==b) // true
    fmt.Println(a=="a") // true
    fmt.Println(exist) // true
}
