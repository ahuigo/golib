package merrors

import (
	"fmt"

	"github.com/pkg/errors"
)

type ErrorType string

type HttpCode int

// 只能是数字
//
//go:generate stringer -type HttpCode -linecomment
const (
	HttpOK  HttpCode = 200 //正常
	Http404 HttpCode = 403 //无权限
)
const (
	NetworkError   ErrorType = "error-network"   //网络错误
	NetworkTimeout ErrorType = "network-timeout" //网络超时
	URLError       ErrorType = "error-url"       //URL错误
)

type Error struct {
	ErrType     ErrorType
	HttpCodeTmp HttpCode
	Err         error
	Data        interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%+v", e.ErrType, e.Err)
}

func New(errType ErrorType, msg string) *Error {
	err := errors.New(msg)
	return &Error{
		ErrType: errType,
		Err:     err,
	}
}

func Errorf(errType ErrorType, format string, args ...interface{}) *Error {
	err := errors.Errorf(format, args...)
	return &Error{
		ErrType: errType,
		Err:     err,
	}
}

func Wrap(errType ErrorType, err error, msg string) *Error {
	err = errors.Wrap(err, msg)
	return &Error{
		ErrType: errType,
		Err:     err,
	}
}

func Wrapf(errType ErrorType, err error, format string, args ...interface{}) *Error {
	err = errors.Wrapf(err, format, args...)
	return &Error{
		ErrType: errType,
		Err:     err,
	}
}
