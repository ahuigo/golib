package main

import (
	"errors"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

type writer struct {
	io.Writer
	timeFormat string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

func TestLogPrefix(t *testing.T) {
	// With a flag
    // 全局 log.SetFlags(log.Lshortfile)
    //logger2 := log.New(os.Stdout, "[info] ", log.Lshortfile) // 增加filepath:lineno
	logger2 := log.New(&writer{os.Stdout, "2006/01/02 15:04:05 "}, "[info] ", log.Lshortfile)
	logger2.Println("Hello world!")
	// 2016/07/14 16:50:31 [info] main.go:28: Hello world!
	err := errors.New("EEEEEEE")
	logger2.Println(err)

	// Without  flag
	logger := log.New(&writer{os.Stdout, "2006/01/02 15:04:05 "}, "[info] ", 0)
	logger.Println("Hello world!")
	// 2016/07/14 16:47:04 [info] Hello world!

}
