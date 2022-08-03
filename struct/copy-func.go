package main
 
import (
	"fmt"
)

type Person struct {
    name    string
}

func rename(p Person) Person {
     p.name = "test"
     return p
}
 
func main() {
     p := Person{"richard"}
     np := rename(p)
     fmt.Println(p.name, np.name) //richard test
}


