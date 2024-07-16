package main
import "fmt"

func main(){
    x := []int{1,2,100}
    y := x[:2:2]
    y[0]=11
    y=append(y, 3) // 与x 分离了
    y[0]=12
    fmt.Printf("%v\n", x) //11,2,100
}
