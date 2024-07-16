package main

import (
	"os"
	"runtime/trace"
)

/*
*
go run main.go 2> trace.out
go tool trace --http=':8080' trace.out
*/
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	ch := make(chan int)

	go func() {
		ch <- 42
	}()
	<-ch
}
