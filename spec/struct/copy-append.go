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
        cat:= Cat{7, "Wilson", []string{"Tom", "Tabata", "Willie"}}
        s := append([]Cat{}, cat)
        s[0].age=6
        s[0].friends[0]="newuser"
        fmt.Println(s)
    }

