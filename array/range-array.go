package main
import "fmt"
func main(){
    roles := []string{"hi","le"}
    for i, role:= range roles{
        fmt.Printf("i=%v,v=%v\n", i, role)
        roles[i]="l"
        roles = append(roles, "123")
    }
    for i, role:= range roles{
        fmt.Printf("%v:%v\n", i, role)
    }
}
