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
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	<-ch
}
