// go run -ldflags="-X 'main.Version=v1.0.2'" ldflags-var.go
package main

import (
    "fmt"
)

var Version = "development"

func main() {
    fmt.Println("Version:\t", Version)
}
