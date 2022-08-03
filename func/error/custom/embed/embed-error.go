package merrors

import (
	"fmt"

	"github.com/pkg/errors"
)

type errorHttp struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Err    error  `json:"err"`
}

type ErrorNetwork struct {
	errorHttp
}
type ErrorURL struct {
	errorHttp
}

type ErrorNotImplemented struct {
	errorHttp
}

func (e *errorHttp) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s %s err:%s", e.Method, e.Url, e.Err.Error())
	} else {
		return fmt.Sprintf("%s %s failed", e.Method, e.Url)
	}
}

func NewErrorNetwork(method string, url string, format string, args ...interface{}) *ErrorNetwork {
	err := errors.Errorf(format, args...)
	return &ErrorNetwork{
		errorHttp{
			method,
			url,
			err,
		},
	}
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
