package main

import (
    "fmt"
    "time"
)


/*
可以读写nil chan, 但是会被永远阻塞
*/
func testChan(c chan int){
    go func(){
        println("read chan:",<-c)
    }()
    go func(){
        c<-1
        println("input chan end")
    }()
}
func main(){
    c1 := make(chan int, 0) 
    c2 := make(chan int) 
    var c3 chan int
    fmt.Printf("%#v\n", c1==nil) //false
    fmt.Printf("%#v\n", c2==nil) //false
    fmt.Printf("%#v\n", c2==c1) //false
    fmt.Printf("%#v\n", c3==nil) //true
    //c3 := make(chan int) 
    testChan(c3)

    time.Sleep(2*time.Second)

}
