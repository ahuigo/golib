package main

import (
	"fmt"
	"strings"
)

func main() {
    s:="abcabc2"
    fmt.Println(strings.Split(s,":")[0])
}
