package main

import (
	"fmt"
)
func testNilBytes(){
    var b []byte
    fmt.Printf("b is Nil=%v\n", b==nil) // true
    b=nil
    fmt.Printf("b is Nil=%v\n", b==nil) // true
    b=[]byte(nil)
    fmt.Printf("b is Nil=%v\n", b==nil) // true
    b=[]byte{}
    fmt.Printf("b is not Nil=%v\n", b==nil) // true
    b=nil
    fmt.Printf("nil slice append:%v\n", append(b, 1,2)) // append true
}


func Include[T comparable](array []T, element T) bool {
	for _, elem := range array {
		if elem == element {
			return true
		}
	}
	return false
}

func main() {
    testNilBytes()

    // foreach nil
    var s []int
    println("nil?:", s==nil)
    if Include(s, 1){
        println("exist")
    }else{
        println("not exist")
    }
}
