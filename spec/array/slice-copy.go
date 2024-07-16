package main

import "fmt"

func main() {
    src := []string{"a", "b", "c"}
    dst := make([]string, len(src)-1)

    // dst=append(dst, src...)
    copy(dst, src)

    fmt.Printf("source slice: %[1]v, address: %[1]p\n", src)
    fmt.Printf("source slice: %[1]v, address: %[1]p\n", dst)
}
