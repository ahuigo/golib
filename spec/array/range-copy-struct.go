package main
import "fmt"
type Role struct{
    Users []string
}
type Roles []*Role
func main(){
    roles := Roles{&Role{[]string{"ahui"}}}
    // 如果用非指针，roles就在堆在，role[i]被从堆copy 到栈上(对于大struct 有性能问题)
    // 如果不用pointer, range 就会copy values
    for _, role:= range roles{
        role.Users = append(role.Users, "new value")
        fmt.Println("role:",role)
        fmt.Println("roles[0]:",roles[0])
        fmt.Printf("%p,%p\n", role,roles[0])
    }
}
