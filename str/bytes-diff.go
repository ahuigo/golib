package main

import (
	"fmt"
)

// return string(a)==string(b)
func BytesDiff(a, b []byte) bool {
	al := len(a)
	bl := len(b)
	if al != bl {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

/**
  -1 no diff
  0 diff at begin
  1 diff at 1 byte
    a:=[]byte("abc")
    fmt.Println(BytesDiffn(a,[]byte("abc"))==-1)
    fmt.Println(BytesDiffn(a,[]byte("abcde"))==3)
    fmt.Println(BytesDiffn(a,[]byte("abde"))==2)
    fmt.Println(BytesDiffn(a,[]byte("ab"))==2)
*/
func BytesDiffn(a, b []byte) int {
	if len(a) > len(b) {
        a,b=b,a
	}
    i,v:=0,byte(0)
	for i, v = range a {
		if v != b[i] {
			return i
		}
	}
    if len(a)==len(b){
        return -1
    }else{
        return i+1
    }
}


func main() {
    a:=[]byte("abc")
    fmt.Println(BytesDiffn(a,[]byte("abc"))==-1)
    fmt.Println(BytesDiffn(a,[]byte("abcde"))==3)
    fmt.Println(BytesDiffn(a,[]byte("abde"))==2)
    fmt.Println(BytesDiffn(a,[]byte("ab"))==2)
}

