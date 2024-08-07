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
    cat1 := &Cat{7, "cat1", []string{"Tom", "Tabata", "Willie"}}
    cat2 := new(Cat)
    cat2.name = "cat2"
    fmt.Println(cat1)
    fmt.Println(cat2)

    *cat2 = *cat1
    cat1.friends[0] = "Tom1"
    fmt.Println(cat1) // &{7 cat1 [Tom1 Tabata Willie]}
    fmt.Println(cat2) //&{7 cat1 [Tom1 Tabata Willie]}

}
