//链接：https://www.zhihu.com/question/425625461/answer/2356045938

package demo

import (
	"os"
	"testing"
)

func returnNamedValue1() (err error) {
	if fh, err := os.Open("abc"); err != nil {
		println("open failed: ", err.Error())
		return err //必须
	} else {
		_ = fh
		println("open success")
	}
	return
}
func returnNamedValue2() (err error) {
	fh, err := os.Open("abc")
	_ = fh
	return
}

func TestReturnNamedValue(t *testing.T) {
	println(returnNamedValue1() != nil) // true
	println(returnNamedValue2() != nil) // true
}
