package main
import "fmt"

func main(){
    x := []int{1,2}
    y := x[:2]
    y=append(y, 3)
    x[0]=11
    fmt.Printf("%v\n", y) //1,2,3
}
