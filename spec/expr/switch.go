package main

import (
	"fmt"
)

func main() {
    status := 1+2
    a, b := false, true
    switch true{
    case status==1, a, b:
        fmt.Println(status)
        fallthrough
    case status==2:
        fmt.Println("fallthrough")
    default:
        fmt.Println("return") 
        return
    }
}
