package main

/*
缓冲的channel:
https://aceld.gitbooks.io/how-do-go/content/channel/dai-huancun-de-channel.html
*/

import (
        "fmt"
)

func main() {
    ch:= make(chan int, 10)
    ch<-1
    ch<-1
    fmt.Println("len:",len(ch)) // "2" 查看channel有效数据
    fmt.Println("cap:",cap(ch)) // "10" 查看channel容量

}
