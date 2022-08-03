package main

import (
    "fmt"
)

func main() {
    var i interface{} = nil
    fmt.Printf("obj=%v,%T\n", i,i)
    fmt.Printf("obj=%v,%T\n", i==nil,i==nil)
}
