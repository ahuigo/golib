package demo

import (
	"fmt"
	"testing"
    "context"
)

func TestInterfaceNil(t *testing.T) {
	var i interface{} = nil
	fmt.Printf("obj=%v,%T\n", i, i) // nil,nil

    var a context.Context
    var b any = a
    //b = context.Background()
    if _, ok := b.(context.Context); ok{
        fmt.Println("nil is context(wrong)")
    }else{

        fmt.Printf("nil is not context,Type=%T\n", b)//打印: nil is not context,Type=<nil>
    }
}
