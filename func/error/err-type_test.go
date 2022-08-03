package demo

import (
	"errors"
	"testing"
	// "github.com/pkg/errors"
)

type NetworkError struct {
	Err error
}

type NetworkTimeoutError struct {
	Err error
}

// 服务内部错误
type ServiceError struct {
	Err error
}

func (e *ServiceError) Error() string {
	return e.Err.Error()
}

func (e *NetworkError) Error() string {
	return e.Err.Error()
}

func Test(t *testing.T) {
	var err error = &NetworkError{}

	if _, ok := err.(*NetworkError); !ok {
		t.Fatalf("not network error")
	}

	target := &NetworkError{}
	if !errors.As(err, &target) {
		t.Fatalf("errors.As: not network error")
	}

	// if !errors.Is(err, &NetworkError{}) // NetworkError.Is()

}
