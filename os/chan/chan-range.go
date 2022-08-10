package main

import "fmt"

//https://aceld.gitbooks.io/how-do-go/content/channel/chuan-lian-de-channel-pipeline.html
func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}
