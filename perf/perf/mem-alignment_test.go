package perf

import (
	"fmt"
	"testing"
	"unsafe"
)

// 内存对齐　https://geektutu.com/post/hpg-struct-alignment.html
/*
*

对齐作用：避免数据跨越边界字长时，处理器可能需要进行额外的内存访问操作
对齐保证(align guarantee), 对齐倍数：

	对于任意类型的变量 x ，unsafe.Alignof(x) 至少为 1。
	对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的 unsafe.Alignof(x.f)，unsafe.Alignof(x) 等于其中的最大值。
	对于 array 数组类型的变量 x，unsafe.Alignof(x) 等于构成数组的元素类型的对齐倍数。
*/
func TestArgAlign(t *testing.T) {
	type Flag struct {
		Num1 int16
		Num2 int32
	}
	fmt.Println(unsafe.Alignof(Flag{}))                 // 4
	fmt.Println(unsafe.Alignof(struct{ Num1 int16 }{})) // 2
	s := ""
	fmt.Println(unsafe.Alignof(s)) // 8
}

// 如果变量的大小不超过字长，那么内存对齐后，对该变量的访问就是原子的. 高效＋减少并发问题
func TestArgSize(t *testing.T) {
	type demo1 struct {
		a int8  // 1
		b int16 // 2, 对齐倍数为 2，因此，必须空出 1 个字节，偏移量才是 2 的倍数
		c int32 // 4
	}

	type demo2 struct {
		a int8  // 1
		c int32 // 4, 对齐倍数为 4，因此，必须空出 3 个字节，偏移量才是 4 的倍数
		b int16 // 2
	} // 整个结构体对齐倍数为 4，因此，必须再空出 2 个字节，偏移量才是 4 的倍数: 1+(3)+4+2+(2)=12

	fmt.Println(unsafe.Sizeof(demo1{})) // 8
	fmt.Println(unsafe.Sizeof(demo2{})) // 12

	// struct{} 大小为 0，作为其他 struct 的字段时，一般不需要内存对齐。但是有一种情况除外：即当 struct{} 作为结构体最后一个字段时，需要内存对齐。因为如果有指针指向该字段, 返回的地址将在结构体之外，如果此指针一直存活不释放对应的内存，就会有内存泄露的问题
	type demo3 struct {
		c int32    // 4
		a struct{} // 0(+4)
	} // size: 4+4=8

	type demo4 struct {
		a struct{}
		c int32 //4
	}

	fmt.Println(unsafe.Sizeof(demo3{})) // 8
	fmt.Println(unsafe.Sizeof(demo4{})) // 4

}
