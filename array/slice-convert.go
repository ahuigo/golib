package main
import (
    "fmt"
    "unsafe"
)

type Int int
type Ints []int

func type_convert_safe(){
    var c = Ints{1, 2}

    // 因为[]int 与 Ints 虽然是不同的type, 但是有相同的interface method
    var x []int = c
    c = x
    fmt.Println(x)
}

func type_convert_safe2(){
    var c = []int{1, 2}

    var x Ints = c
    fmt.Println(x)
}

func type_convert_unsafe(){
    var c = Ints{1, 2}

    var x []int = *(*[]int)(unsafe.Pointer(&c))
    fmt.Println(x)
}


func alias_convert_safe(){
    var c = []Int{1, 2}
    x := make([]int, len(c))
    for i, v := range c {
        x[i] = int(v)
    }
    fmt.Println(x)
}


func alias_convert_unsafe(){
    var c = []Int{1, 2}

    var x []int = *(*[]int)(unsafe.Pointer(&c))
    fmt.Println(x)
}







func main(){
    type_convert_safe()
    type_convert_safe2()
    type_convert_unsafe()
    alias_convert_safe()
    alias_convert_unsafe()
}
