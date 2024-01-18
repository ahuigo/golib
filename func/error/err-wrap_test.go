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
	err = fmt.Errorf("wrap with func:%w", err)

	fmt.Printf("error:\n%+v\n", err)
	// Example output:
	// wrap with func:whoops
}
