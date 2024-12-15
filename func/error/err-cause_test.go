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
	err = errors.Wrapf(err, "main(wrap3)")
	if err != nil {
		fmt.Printf("original (%T): %v\n\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("error():%v\n\n", err.Error())
		fmt.Printf("stack trace:\n %+v\n\n", err)                                             // %+v 可以在打印的时候打印完整的堆栈信息
		fmt.Printf("is PathError As: %+v\n", isPathError(err))                                // true
		fmt.Printf("is PathError Strict: %+v\n", isPathErrorStrict(err))                      // false
		fmt.Printf("is PathError Strict(Cause): %+v\n", isPathErrorStrict(errors.Cause(err))) // true
		// os.Exit(1)
	}
}

func isPathError(err error) bool {
	var target *os.PathError
	// return errors.Is(err, target) // 不能用Is, 因为Is是递归判断`error值``, As是for递归As/Unwrap判断`error类型``
	return errors.As(err, &target) // 会改写target
}

func isPathErrorStrict(err error) bool {
	_, ok := err.(*os.PathError)
	return ok
}
