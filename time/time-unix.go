package main

import (
    "fmt"
    "time"
)

func main() {
    //1626177311151735000
    fmt.Printf("%T\n", 1626177311151735000/1e6)
    i := 1626177311991735000/1e9
    fmt.Printf("%T, i=%v\n", i,i)
    seconds:= int64(i)
    fmt.Printf("seconds=%#v\n", seconds)
    fmt.Println(time.Unix(seconds, 0))

}
