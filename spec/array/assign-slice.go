package main
import "fmt"
func main(){
    var ns []int
    ns=append(ns,1)
    fmt.Printf("nil slice?%v\n", ns)  //same

    a := []int{1,2,3}
    b := a
    c :=a[0:0]
    c = append(c, 50,51,52)
    a[1]=100
    fmt.Printf("%v,%v,%v\n", a,b,c)  //same
    fmt.Printf("%p,%p,%p\n", a,b,c)  //same
}
