package main
import "fmt"
import "github.com/thoas/go-funk"

func main(){
    a:=[]string{"a","d"}
    b:=[]string{"a","b","c"}
    fmt.Println(b)

    res := funk.Contains(a,b) //false
    fmt.Println(res)
    res = funk.Subset(a,b) //true
    fmt.Println(res)
}
