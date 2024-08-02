package demo

import (
	"bufio"

	"github.com/pkg/errors"
)

func TestBufio() {
	reader := bufio.NewReader(os.Stdin)
    line, err = reader.ReadString('\n')
    println(line)
}
