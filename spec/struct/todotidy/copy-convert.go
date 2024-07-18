package main
 
import (
	"fmt"
)
 
    type Cat struct {
        age     int
        name    string
    }
    type Cat2 struct {
        age     int
        name    string
    }

    func main() {
        // copy struct mehtod: 
        //1. sp:=s
        //1. sp:=Type(s)
        cat := Cat{age:7}
        cat2 := Cat2(cat)
        cat2.age=1

        fmt.Println(cat)
        fmt.Println(cat2)
    }
