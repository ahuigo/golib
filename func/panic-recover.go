package main
import "fmt"
func main(){
    defer func() {
        r:=recover()
        fmt.Printf("Recovered return: %#v, type:%T\n",r, r)
        r=recover()
        fmt.Printf("Recovered return: %#v, type:%T\n",r, r)
    }()
    panic([]string{"panic error!"})
}
