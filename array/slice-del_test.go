package main

import "fmt"
import "testing"

func TestSliceDel(t *testing.T) {
/*
    a=append(a[:i],a[i+1:]...)
    # or i+1 往前复制
    a=a[:i+copy(a[i:],a[i+1:])]

删除后设置为nil 方便内存回收（好像不智能？）

    if i<len(a)-1{
        copy(a[i:],a[i+1:])
    }
    a[len(a)-1]=nil
    a=a[:len(a)-1]
    */

    a := []string{"a", "b", "c"}
    i:=1
    a=a[:i+copy(a[i:],a[i+1:])]

    fmt.Printf("source slice: %[1]v, address: %[1]p\n", a)
}
