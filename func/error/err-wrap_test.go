package demo

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestWrap2(t *testing.T) {
	f := func() error {
		// 没有%w, 等价于errors.New("whoops")
		err := fmt.Errorf("whoops")
		return err
	}
	err := f()
	// 有%w, 返回wrapError/wrapError2 实例
	err = fmt.Errorf("wrap with errorf %%w:%w", err)

	fmt.Printf("error: %+v\n", err)
	fmt.Printf("err.Error(): %s\n\n", err.Error())
	fmt.Printf("errors.Unwrap(err): %v\n", errors.Unwrap(err))
	// Example output:
	// wrap with func:whoops
}
