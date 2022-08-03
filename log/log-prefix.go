package main 

import (
	"io"
	"log"
	"os"
	"time"
)

type writer struct {
	io.Writer
	timeFormat string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

func main(){
	// Without any flags
	logger := log.New(&writer{os.Stdout, "2006/01/02 15:04:05 "}, "[info] ", 0)
	logger.Println("Hello world!")
	// 2016/07/14 16:47:04 [info] Hello world!
	
	
	// With a flag
	logger2 := log.New(&writer{os.Stdout, "2006/01/02 15:04:05 "}, "[info] ", log.Lshortfile)
	logger2.Println("Hello world!")
	// 2016/07/14 16:50:31 [info] main.go:28: Hello world!
}
