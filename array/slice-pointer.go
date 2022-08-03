package main

import "fmt"

func main(){
    var a = []int{1,2,3}
    func (l *[]int){
        *l=(*l)[1:]
    }(&a)
    fmt.Println(a)
}
