package main

import (
    "fmt"
    "regexp"
)

func main() {
    value := "catXdogYbird"

    // First compile the delimiter expression.
    re := regexp.MustCompile(`[XY]`)

    // Split based on pattern.
    // ... Second argument means "no limit."
    result := re.Split(value, -1)

    for i := range(result) {
        fmt.Println(result[i])
    }
}
