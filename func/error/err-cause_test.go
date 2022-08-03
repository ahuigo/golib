package demo

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	_, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed(wrap1)")
	}
	return nil, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.WithMessage(err, "cound not read config(wrap2)")
}

func TestErrCause(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		fmt.Printf("original (%T)err: %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("err:%v\n", err)
		fmt.Printf("stack trace:\n %+v\n", err) // %+v 可以在打印的时候打印完整的堆栈信息
		// os.Exit(1)
	}
}
