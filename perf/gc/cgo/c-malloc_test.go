package main

// 错误的做法 ❌
var cPointer *SomeGoStruct // C 代码不能存储go pointeer

// 正确的做法 ✅
goSlice := []byte{1, 2, 3}

// 1. 在 C 中分配内存
cBuf := C.malloc(C.size_t(len(goSlice)))
defer C.free(cBuf)

// 2. 将 Go 数据拷贝到 C 内存中
// 把 cBuf（unsafe.Pointer 或 *C.void）转换成指向一个非常大的字节数组的指针。
// 这里用 1 << 30（1 GiB）只是为了让数组足够大，不会在后续切片时越界
cSlice := (*[1 << 30]byte)(cBuf)[:len(goSlice):len(goSlice)]
copy(cSlice, goSlice)

// 3. 将 C 的指针 (cBuf) 传递给其他 C 函数
// C.process_data(cBuf, C.int(len(goSlice)))
