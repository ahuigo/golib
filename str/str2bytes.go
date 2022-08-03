package main
/*
作者：Golang梦工厂
链接：https://zhuanlan.zhihu.com/p/386676358
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

因为string和slice的结构字段是相似的：type stringStruct struct {
    str unsafe.Pointer
    len int
}
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
唯一不同的就是cap字段，array和str是一致的，len是一致的，所以他们的内存布局上是对齐的，这样我们就可以直接通过unsafe.Pointer进行指针替换。
*/

func stringtoslicebytetmp(s string) []byte {
 str := (*reflect.StringHeader)(unsafe.Pointer(&s))
 ret := reflect.SliceHeader{Data: str.Data, Len: str.Len, Cap: str.Len}
 return *(*[]byte)(unsafe.Pointer(&ret))
}

func main()  {
     str := "hello"
     by := stringtoslicebytetmp(str)
     by[0] = 'H' // error： 因为string是不可变的，被限制了
}
