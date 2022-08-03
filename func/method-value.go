package main

import(
    "fmt"
)


type T struct {
    a int
}
func (tv  T) Mv(a int) int         { return a }  // value receiver
func (tp *T) Mp(f float32) float32 { return f }  // pointer receiver

var t T
var pt *T = &t
func makeT() T{
    return t
}

func main(){
    /*
    f := t.Mv; f(7)   // like t.Mv(7)
    f := t.Mp; f(7)   // like (&t).Mp(7)
    f := pt.Mp; f(7)  // like pt.Mp(7)
    f := pt.Mv; f(7)  // like (*pt).Mv(7) (pt must be valid pointer)
    f := makeT().Mv; f(7)   // like makeT().Mv(7)
    f := makeT().Mp; f(7)   // (&makeT()).Mp(7) . invalid: result of makeT() is not addressable
    */
    f := makeT().Mp;
    fmt.Println(f(7))


}
