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
        wilson := []Cat{{7, "Wilson", []string{"Tom", "Tabata", "Willie"}}}
        nikita := []Cat{{}}
        copy(nikita, wilson)            //Cat是浅复制
        nikita[0].age=6
        nikita[0].friends[0]="newuser"  //friends 同时改变

        fmt.Println(wilson)             //[{7 Wilson [newuser Tabata Willie]}]
        fmt.Println(nikita)             //[{6 Wilson [newuser Tabata Willie]}]
    }
