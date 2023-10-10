package demo

import "testing"

func TestPtrArgs(t *testing.T) {
	fn := func(s *string) {
		t.Log(s == nil) //false
		*s = "hello"
	}
	var s string
	fn(&s)
	t.Log(s)
}
