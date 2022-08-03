package main

import (
	"fmt"
)

func main() {
    status := 1+2
    switch status {
    case 1:
        fmt.Println(status)
    case 2:
        fmt.Println(status) 
    case 3:
    default:
        fmt.Println("return") 
        return
    }
}
