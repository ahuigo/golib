package demo

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func g() error {
	err := errors.New("whoops")
	return err
}
func f() error {
	err := g()
	err = errors.Wrapf(err, "wrap with func:%s", "f()")
	return err
}

func TestTrace(t *testing.T) {
	err := f()
	fmt.Printf("err.(stackTracer)---------------------------\n")
	if err, ok := err.(stackTracer); ok {
		for _, e := range err.StackTrace() {
			fmt.Printf("%+s:%d(%T)\n", e, e, e)
		}
	}

	fmt.Printf("\nerr.(stackTracer) with cause(original)---------------------------\n")
	if errors.Cause(err) != err {
		if oerr, ok := errors.Cause(err).(stackTracer); ok {
			for _, f := range oerr.StackTrace() {
				fmt.Printf("%+s:%d(%T)\n", f, f, f)
			}
		}
	}

	fmt.Printf("\n\ntrace all--------------\n%+v", err)
	// Example output:
	// whoops
	// github.com/pkg/errors_test.ExampleNew_printf
	//         /home/dfc/src/github.com/pkg/errors/example_test.go:17
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
}
