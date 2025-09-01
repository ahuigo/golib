package main


type MyObject struct {
	// 这个指针指向 C 分配的内存
	cData unsafe.Pointer 
	name  string
}

func NewMyObject() *MyObject {
	obj := &MyObject{
		cData: C.malloc(1024), // 分配了 C 内存
		name:  "my-object",
	}

    // ✅ 设置终结器 指定obj 的回收处理handler
	runtime.SetFinalizer(obj, func(o *MyObject) {
		fmt.Println("Finalizing and freeing C data for", o.name)
		C.free(o.cData)
	})
	return obj
}

func useObject() {
	obj := NewMyObject()
	// ... 使用 obj ...
    _ = obj

	// 当 useObject 函数结束，obj 成为垃圾
	// 如果不加SetFinalizer， Go GC 会回收 obj，但不会释放 obj.cData 指向的内存 -> 内存泄漏！ 
}
