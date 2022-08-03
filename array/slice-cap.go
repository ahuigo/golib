package main
import "fmt"

func main(){
    var a = [10]int{1,2,3,4}

    //# len(b)=cap(b)=10
    b := a[:];  
    fmt.Println(b, len(b), cap(b))

    s := append(b,4,5);  
    fmt.Println(s, len(s), cap(s)) 

    // 
    length := 1
    c1 := make([]int,length,10)
    c := c1[length:]
    fmt.Println(c, len(c), cap(c), ) 
}
