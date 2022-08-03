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
    fmt.Printf("nil slice append:%v\n", append(b, 1,2)) // true
}

func main() {
    testNilBytes()
}
