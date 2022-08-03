package main

import (
        "fmt"
)

func main() {
    // & | ^(XOR) &^(NOT)
    //var bits uint8 = 0x80
    fmt.Printf("0x80=%#b\n",uint8(0x80))
    fmt.Printf("0x80=%#b\n",(uint8(127)^(uint8(3))))
}
