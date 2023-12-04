package ptr

func getPointerAddress(refV reflect.Value) uintptr {
    address:=refV.Pointer()

    // ptr string
    addressStr := fmt.Sprintf("%p", address)

    // 获取指针的地址
    pointer := unsafe.Pointer(address)

    // 将指针地址转换为uintptr类型
    address := uintptr(pointer)

    return address
}
