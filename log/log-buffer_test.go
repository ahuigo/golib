package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestLoggerBuffer(t *testing.T) {
	var buf bytes.Buffer
	var logger = log.New(&buf, "prefix: ", log.Lshortfile)
	logger.Print("log with file:lineno!") // 写入buf
	fmt.Print(&buf)                       //输出buf
}
