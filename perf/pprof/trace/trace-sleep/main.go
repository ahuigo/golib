package main

import (
	"os"
	"runtime/trace"
	"time"
)

/*
*
go run main.go 2> trace.out
go tool trace --http=':8080' trace.out
*/
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	// go func() {
	// 	time.Sleep(time.Microsecond * 50)
	// }()
	// time.Sleep(time.Microsecond * 50)
	time.Sleep(time.Microsecond * 200)
	time.Sleep(time.Microsecond * 200)
	time.Sleep(time.Microsecond * 200)
}
