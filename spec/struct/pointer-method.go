package main
 
import (
	"fmt"
)

type Person struct {
    name    string
}

func (p *Person) rename() {
     p.name = "test"
}
 
func main() {
     p := Person{"richard"}
     p.rename() // use pointer
     fmt.Println(p.name) //test
}


