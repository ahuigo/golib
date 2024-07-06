package demo

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestWriterCopy(t *testing.T) {
    str := "Message\n"
    fmt.Fprintln(os.Stderr, str)
    io.WriteString(os.Stderr, str)
    io.Copy(os.Stderr, bytes.NewBufferString(str))
    os.Stderr.Write([]byte(str))
}
