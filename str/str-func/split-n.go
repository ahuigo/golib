package main

import (
    "fmt"
    "strings"
)

func main() {
    value := "12|34|56|78"
    // Split into three parts.
    // ... The last separator is not split.
    result := strings.SplitN(value, "|", 3)
    for v := range(result) {
        fmt.Println(result[v])
    }
}
