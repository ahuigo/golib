package main
 
import (
	"fmt"
)
 
type Cat struct {
    age     int
    name    string
    friends []string
}

func main() {
    // test1
    cat1 := Cat{7, "cat1", []string{"Tom", "Tabata", "Willie"}}
    cat2 := cat1        //结构体/array会按值，slice按引用传值
    cat2.name = "cat2"
    cat2.friends[0]= "newname"

    fmt.Println(cat1)
    fmt.Println(cat2)

    // test2
    s1:=[]string{"Tom", "Tabata", "Willie"}
    s2:=s1
    s2[0]="Tom2"
    fmt.Println(s1)
    fmt.Println(s2)
}
