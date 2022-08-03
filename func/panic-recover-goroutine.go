package main
import (
    "fmt"
    "time"
)

func g(){
    panic([]string{"some error"})
}

func request(){
    defer func() {
        r:=recover()
        fmt.Printf("Recovered return: %#v, type:%T\n",r, r)
    }()
    go g()
    time.Sleep(time.Second)
    println("response!")
}

func main(){
    go request()
    time.Sleep(time.Hour)
}
