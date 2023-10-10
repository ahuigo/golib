package main
import "fmt"

func main(){
    // define nil slice
    var x2 []int
    fmt.Printf("%#v\n", x2) //[]int(nil)
    fmt.Printf("%#v\n", x2==nil) //true

    // nil slice 的len/cap = 0, 且可以append
    var s []int
    fmt.Println(s, len(s), cap(s)) //[] 0 0
    s = append(s, 1)
    fmt.Println(s, len(s), cap(s)) //[] 0 0
}
