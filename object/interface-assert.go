package main

import (
	"fmt"
)

func main() {
    var i interface{} = "string"
    bytes := []byte{}
    switch v := i.(type) {
        case string:
            bytes = []byte(v)
            fmt.Printf("bytes: %v, type of v: %T", bytes, v)
        default:
            fmt.Printf("I don't know about type %T!\n", v)
    }
}
