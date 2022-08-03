package main

import "fmt"
type Stu struct{
    Name string
}

func main() {
    stu := Stu{"ahui1"}
    stus := []Stu{stu} // this is copy
    stu.Name = "ahui2"
    fmt.Printf("stu: %v\n", stu)
    fmt.Printf("stus: %[1]v\n", stus)
}
