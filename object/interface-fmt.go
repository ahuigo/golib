package main

import (
    "fmt"
)

func main() {
    m:= map[string]interface{}{"a":1}
    fmt.Printf("a=%#T\n", m["a"].(int))
    fmt.Printf("a=%#T\n", m["a"])
    b:= map[string]interface{}{"b":[]string{"str1"}}
    fmt.Printf("b=%#T\n", b["b"])
    // to string

    var n1 interface{} = 1.25
    var n2 interface{} = float64(2222222222222222.22333333333)
    fmt.Printf("n1=%f,%T\n", n1,n1)
    fmt.Printf("n1=%v,%T\n", n1,n1)
    fmt.Printf("n2=%f,%T\n", n2,n2)
    fmt.Printf("n2=%v,%T\n", n2,n2)
    fmt.Printf("obj=%v,%T\n", m,m)
}
