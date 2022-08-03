package main

import "fmt"

func main() {
    key := "age"
    m := map[string]interface{}{
        "a":1,
        "k2":"b",
        key: 20,
    }
    fmt.Printf("%v\n", m)
}
