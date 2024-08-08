package demo

import "testing"

func TestLoopBreak(t *testing.T) {
	i := 1
	foo := 1
	condA := true
	count := 0

loop:
	for {
		count += 1
		switch i {
		case foo:
			if count > 10 {
				break loop
			}
			if condA {
				println("doA")
				break // 'break switch'
			}
			println("doC")
		default:
			if count > 10 {
				break loop
			}
			println("do default")
		}
	}

	println("end")
}
