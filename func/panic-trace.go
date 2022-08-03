package main

func test() {
    panic("test panic2")
}

func parent(){
    test()
}
func main() {
    parent()
}
