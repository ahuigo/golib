package main
func TestGcStr() {
	goString := "hello cgo"
	cString := C.CString(goString)
	defer C.free(unsafe.Pointer(cString)) // ✅ 无论函数如何退出，都会被执行

	// ... 安全地使用 cString ...
}
