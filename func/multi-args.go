package main

import (
	"fmt"
    //"string"
)

func swap(x, y string) (string, string) {
    return y, x
}
func multiArgsHandler(args ...interface{}) () {
    println(args[0].(string), args[1].(string))
}

func main() {
	fmt.Println(swap("a","b")) 
	multiArgsHandler(swap("a","b")) // ok

    //println(swap("a","b"))  //error: multiple-value swap() in single-value context
}

