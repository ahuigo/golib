package main
import "fmt"
func test() {
    defer func() {
        fmt.Println("recover:",recover())
    }()
    defer func() {
        panic("test panic2:"+recover())
    }()
    panic("test panic1")
}

func parent(){
    test()
}
func main() {
    parent()
}
