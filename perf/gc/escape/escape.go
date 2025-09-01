package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// 参考：https://tonybai.com/2021/05/24/understand-go-escape-analysis-by-example/
// go run -gcflags='-m' escape.go
type Person struct {
	Name string
}

// 大结构体 -  测试栈限制
type LargePerson struct {
	Name    string
	Address string
	Email   string
	Phone   string
	Date    [20000]int64 // ~800KB(64位系统上 栈限制一般是800KB)
	// Metadata    map[string]string
}

// 极大的结构体 -
type ExtremelyLargePerson struct {
	Name       string
	HugeArray1 [100000]int64 // ~800KB
	// 总共约2.4MB
}

//go:noinline
func TestLargeString() string {
	s := func() string {
		s := "a" // 2^(20+4) = 16MB
		for i := 0; i < (20 + 4); i++ {
			s += s
		}
		return s
	}()
	return s
}

//go:noinline
func TestLargeArray() int64 {
	// 创建一个包含 81920 个 int 的数组。640KB。超过编译器的栈分配阈值。
	var arr [81920]int64 //moved to heap: arr
	// 使用数组，防止被优化掉
	for i := range arr {
		arr[i] = int64(i)
	}
	return arr[0]
}

func NoEscapePointer() {
	p2 := &Person{Name: "Alice"}
	_ = p2
}

// 返回any 类型会逃逸: p escapes to heap
func EscapeInterface() any {
	p := Person{Name: "Alice"}
	return p
}

// no escape small struct
func NoEscapeSmallStruct() Person {
	p := Person{Name: "Bob"}
	return p // small结构体不逃逸
}

// 大结构体是否会逃逸:moved to heap: large
func EscapeLargeStruct() LargePerson {
	large := LargePerson{ // 大结构体一生成时就逃逸了: escapes to heap
		Name:    "Large Bob",
		Address: "123 Big Street",
		Email:   "large@example.com",
	}
	return large
}

// 测试极大结构体 // moved to heap: extreme
func EscapeExtremelyLargeStruct() ExtremelyLargePerson {
	extreme := ExtremelyLargePerson{ // 极大结构体生成时就逃逸了: escapes to heap
		Name: "Extreme Bob",
	}
	return extreme
}

// map[string]int{...} escapes to heap
func EscapeMap() map[string]int {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	return m // escapes to heap
}

func TestUnsafePointer() unsafe.Pointer {
	v := 42
	p := unsafe.Pointer(&v) // v逃逸
	return p
}
func TestUnsafePointer2() unsafe.Pointer {
	p := TestUnsafePointer() // 这里 p 逃逸
	return p
}

// 获取变量内存大小的不同方法
func GetVariableSize() {
	fmt.Println("\n=== 变量内存大小测试 ===")

	// 1. 使用 unsafe.Sizeof() - 获取类型的字节大小
	var person Person
	var largePerson LargePerson
	var extremelyLarge ExtremelyLargePerson

	ls := TestLargeString()
	fmt.Println(len(ls)) // 使用 ls
	fmt.Printf("large string 大小: %d bytes\n", len(ls))
	fmt.Printf("Person 结构体大小: %d bytes\n", unsafe.Sizeof(person))
	fmt.Printf("LargePerson 结构体大小: %d bytes (%.2f KB)\n",
		unsafe.Sizeof(largePerson), float64(unsafe.Sizeof(largePerson))/1024)
	fmt.Printf("ExtremelyLargePerson 结构体大小: %d bytes (%.2f MB)\n",
		unsafe.Sizeof(extremelyLarge), float64(unsafe.Sizeof(extremelyLarge))/1024/1024)

	// 2. 使用 reflect.TypeOf().Size() - 反射方式
	fmt.Printf("Person (reflect): %d bytes\n", reflect.TypeOf(person).Size())
	fmt.Printf("LargePerson (reflect): %d bytes\n", reflect.TypeOf(largePerson).Size())

	// 3. 测试不同基本类型的大小
	fmt.Println("\n基本类型大小:")
	fmt.Printf("int8: %d bytes\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int64: %d bytes\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("string: %d bytes\n", unsafe.Sizeof(""))
	fmt.Printf("[]int: %d bytes\n", unsafe.Sizeof([]int{}))
	fmt.Printf("map[string]int: %d bytes\n", unsafe.Sizeof(map[string]int{}))

	// 4. 实际创建的变量
	largeInstance := EscapeLargeStruct()
	fmt.Printf("\n实际创建的 LargePerson 实例大小: %d bytes (%.2f KB)\n",
		unsafe.Sizeof(largeInstance), float64(unsafe.Sizeof(largeInstance))/1024)

	// 使用实例避免未使用警告
	fmt.Printf("实例名称: %s\n", largeInstance.Name)
}

// 监控内存使用情况
func MonitorMemoryUsage(name string) {
	var m runtime.MemStats
	runtime.GC() // 强制GC
	runtime.ReadMemStats(&m)

	fmt.Printf("\n=== %s 内存统计(heap/stack 以及runtime gcMeta/goroutine/type 等内存占用) ===\n", name)
	fmt.Printf("当前分配内存: %d KB\n", bToKb(m.Alloc))
	fmt.Printf("累计分配: %d KB\n", bToKb(m.TotalAlloc))
	fmt.Printf("go系统申请的内存: %d KB\n", bToKb(m.Sys))
	fmt.Printf("GC次数: %d\n", m.NumGC)
}

func bToKb(b uint64) uint64 {
	return b / 1024
}

func TestEscape() {
	fmt.Println("Testing escape analysis...")

	// 监控初始内存
	MonitorMemoryUsage("程序开始")

	// 获取变量大小信息
	GetVariableSize()

	// 创建大结构体并显示其大小
	large := EscapeLargeStruct()
	fmt.Printf("\nLarge struct 大小: %d bytes (%.2f KB)\n",
		unsafe.Sizeof(large), float64(unsafe.Sizeof(large))/1024)
	fmt.Printf("Large struct 名称: %s\n", large.Name)

	// 监控创建大结构体后的内存
	MonitorMemoryUsage("创建大结构体后")

	// 创建极大结构体
	extreme := EscapeExtremelyLargeStruct()
	fmt.Printf("\nExtreme struct 大小: %d bytes (%.2f MB)\n",
		unsafe.Sizeof(extreme), float64(unsafe.Sizeof(extreme))/1024/1024)
	_ = extreme

	// 监控最终内存
	MonitorMemoryUsage("创建极大结构体后")
}
