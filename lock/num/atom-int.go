package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func main() {
    var u uint64

    for i := 0; i < 400; i += 1 {
        go func() {
            atomic.AddUint64(&u, 1)
            //u = u+1
        }()
    }

    time.Sleep(time.Second)
    fmt.Println(atomic.LoadUint64(&u))
}
