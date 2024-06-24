package demo

import (
	"fmt"
	"testing"
)

func TestWrap2(t *testing.T) {
	f := func() error {
		err := fmt.Errorf("whoops")
		return err
	}
	err := f()
	err = fmt.Errorf("wrap with errorf %%w:%w", err)

	fmt.Printf("error:%+v\n", err)
	fmt.Printf("err.Error():%s\n", err.Error())
	// Example output:
	// wrap with func:whoops
}
