package goid

import goid1 "github.com/petermattis/goid"
func Get() int64 {
    // goroutine id
	return goid1.Get()
}
