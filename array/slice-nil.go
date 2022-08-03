package main
import "fmt"

func main(){
    var ages []int
    stus:=[]int{}
    fmt.Printf("[]int(nil) : %#v, isNil:%v\n", ages, ages==nil)
    fmt.Printf("[]int{} :%#v\n", stus)
    ages = append(ages, 1)
    fmt.Printf("ages :%#v\n", ages)
}
