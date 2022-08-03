package main

import (
    "math/rand"
    "fmt"
    "testing"
)


func StringSliceShuffle(src []string) []string {
	dest := make([]string, len(src))
	ints := rand.Perm(len(src))
    fmt.Println(ints)
	for idx, elem := range ints {
		dest[elem] = src[idx]
	}
	return dest
}

func TestRandPerm(t *testing.T){
    s:=[]string{"a","b","c","d","e"}
    fmt.Println(StringSliceShuffle(s))
}
