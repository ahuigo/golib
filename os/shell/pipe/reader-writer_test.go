package demo

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestStringReader(t *testing.T) {
	sr := strings.NewReader("hi, alex")
	b := make([]byte,2)
	if n, err := sr.Read(b); err != nil {
		t.Fatal(err)
	} else {
		t.Log("read string reader:", string(b[:n])) // n=2
	}

}

func TestStringBuilder(t *testing.T) {
	var sb strings.Builder
	sb.Write([]byte("hi"))
	sb.WriteByte(',')
	sb.WriteString(" Alex!")
	t.Log(sb.String())
}

func TestBytesBuffer(t *testing.T) {
	// reader + writer
	// var sr *bytes.Buffer
	// 1. use String()
	sr := bytes.NewBufferString("hi")
	sr.Write([]byte(", Alex!"))
	t.Log(sr.String())

	// 2. use Read()
	b := make([]byte, 2)
	if n, err := sr.Read(b); err != nil {
		t.Fatal(err)
	} else {
		t.Log("read bytes buffer:", string(b[:n])) // n=2
	}
}

func TestReaderAsStdoutStderr(t *testing.T) {
    var stderr bytes.Buffer

    cmd := exec.Command("cat")
    cmd.Stdin = bytes.NewBuffer([]byte("hi"))
    cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr

    cmd.Run()

}