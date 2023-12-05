package ptr

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestPointerAddress(t *testing.T) {
	var age int = 10
	agei := &age
	refV := reflect.ValueOf(agei)
	address := refV.Pointer()

	fmt.Printf("address uintptr:%v\n", address)

	// 获取指针的地址
	pointer := unsafe.Pointer(address)

	// 将指针地址转换为uintptr类型
	intptr := uintptr(pointer)
	fmt.Printf("address uintptr:%v\n", intptr)

}
