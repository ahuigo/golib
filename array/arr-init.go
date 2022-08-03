package main
import "fmt"

func main(){
    x := []int{1,2}
    fmt.Printf("%v\n", x) //1,2

    // nil
    var x2 []int
    fmt.Printf("%#v\n", x2) //[]int(nil)
    fmt.Printf("%#v\n", x2==nil) //true

    // nil/非nil的len/cap相同
    var s []int
    fmt.Println(s, len(s), cap(s)) //[] 0 0
    s=[]int{}
    fmt.Println(s, len(s), cap(s)) //[] 0 0
}
