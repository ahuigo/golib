package main

import (
    "fmt"
    "time"
)

func main() {
    ztimer := time.NewTimer(0)
	for {
		select {
		case <-ztimer.C:
            //ztimer.Reset(0 * time.Millisecond)
            fmt.Println("haha")
		}
	}
}
