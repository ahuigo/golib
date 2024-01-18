package demo

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func f1() error {
	err := errors.New("whoops")
	return err
}
func f2() error {
	err := f1()
	err = errors.Wrapf(err, "wrap with func:%s", "f()")
	return err
}

func TestTrace(t *testing.T) {
	err := f2()
	fmt.Printf("1. err.(stackTracer): empty original error---------------------------\n")
	if err, ok := err.(stackTracer); ok {
		for _, e := range err.StackTrace() {
			fmt.Printf("%+s:%d(%T)\n", e, e, e)
		}
	}

	fmt.Printf("\n2. err.(stackTracer): include cause(original)---------------------------\n")
	if errors.Cause(err) != err {
		if oerr, ok := errors.Cause(err).(stackTracer); ok {
			for _, f := range oerr.StackTrace() {
				fmt.Printf("%+s:%d(%T)\n", f, f, f)
			}
		}
	}

	fmt.Printf("\n\n3. trace all--------------\n%+v", err)

	err2 := fmt.Errorf("wrap err2:%w", err)
	fmt.Printf("\n\n4.fmt.wrap %%w err2-------------:\n%+v\n", errors.Unwrap(err2))
}
