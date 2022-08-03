package main

import (
	"fmt"
    "time"
)

func main() {
    dst := make(chan int, 0)
    n := 1
    go func() {
        for{
            select {
                case dst <- n:
                    fmt.Println("send n:",n)
                    n++
            }
        }
	}()

    fmt.Printf("start:\n")
    time.Sleep(3*time.Second)
    fmt.Printf("dst:%v\n", <-dst)
    time.Sleep(3*time.Second)
    fmt.Printf("dst:%v\n", <-dst)
}
