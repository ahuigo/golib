package main

import(
    "fmt"
)

// refer: go-var-addressable.md

func main(){
    m:=map[string]string{
        "k1":"val1",
    }
    s:=[2]int{1,2}
    fmt.Println(m["k1"])
    fmt.Println(&s[0])
    fmt.Printf("%v\n",&[]int{123}[0])

}
