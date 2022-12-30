/*
golang里通过mmap拿到的内存，再通过unsafe转换成struct，应该不会被gc吧？ - 重归混沌的回答 - 知乎
https://www.zhihu.com/question/264772353/answer/2776158164
*/
package main

import (
    "os"
    "unsafe"
)

type Foo struct {
    a uint64
}
var pool []uint64
//go:noinline
func Alloc(direct bool) *Foo {
    if direct {
        return &Foo{}
    } else {
        var f *Foo
        size := unsafe.Sizeof(*f)
        need := (size + 7) / 8
        if len(pool) < int(need) {
            pool = make([]uint64, 512)
        }
        p := unsafe.Pointer(&pool[0])
        pool = pool[need:]
        return (*Foo)(p)
    }
}

func main() {
    direct := len(os.Args) == 1
    for i := 0; i < 64*1024*1024; i++ {
        Alloc(direct)
    }
}
/*
$ time ./a
./a  0.77s user 0.01s system 110% cpu 0.710 total
$ time ./a x
./a x  0.26s user 0.06s system 129% cpu 0.250 total
*/
