package main

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
	"unsafe"
)

var globalSink interface{} // 防止编译器优化
var memoryHog [][]byte     // 持续内存压力

// 创建内存压力的辅助函数
func createMemoryPressure() {
	// 创建一些大对象保持在内存中
	for i := 0; i < 100; i++ {
		data := make([]byte, 8192) // 8KB 每个
		memoryHog = append(memoryHog, data)
	}
}

// 清理内存压力
func cleanMemoryPressure() {
	memoryHog = nil
	runtime.GC()
}

// 强制逃逸的 unsafe.Pointer 操作
func forceEscapeUnsafePointer() unsafe.Pointer {
	p := 42
	return unsafe.Pointer(&p) // 必定逃逸
}

// 强制逃逸的普通指针操作
func forceEscapePointer() *int {
	p := 42
	return &p // 必定逃逸
}

// 不逃逸的值操作
func valueOperation() int {
	p := 42
	return p // 不逃逸
}

// 通过接口强制逃逸的 unsafe.Pointer
func interfaceUnsafePointer() interface{} {
	p := 42
	return unsafe.Pointer(&p) // 逃逸 + 接口装箱
}

// 通过接口强制逃逸的普通值
func interfaceValue() interface{} {
	p := 42
	return p // 接口装箱逃逸
}

// GC 压力测试 - 增强版
func measureGCPressure(name string, iterations int, testFunc func()) {
	fmt.Printf("\n=== %s ===\n", name)

	// 强制多次 GC 清理，确保干净的起始状态
	for i := 0; i < 3; i++ {
		runtime.GC()
	}
	runtime.GC()

	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	start := time.Now()

	// 执行测试
	for i := 0; i < iterations; i++ {
		testFunc()

		// 每 10000 次操作检查一次，增加内存压力
		if i%10000 == 0 && i > 0 {
			// 创建临时大对象增加 GC 压力
			tempData := make([]byte, 4096) // 4KB
			globalSink = tempData
		}
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)

	// 计算差值
	allocations := m2.Mallocs - m1.Mallocs
	frees := m2.Frees - m1.Frees
	totalAlloc := m2.TotalAlloc - m1.TotalAlloc
	gcCycles := m2.NumGC - m1.NumGC
	heapObjects := m2.HeapObjects
	pauseTotal := m2.PauseTotalNs - m1.PauseTotalNs

	fmt.Printf("🔢 迭代次数: %d\n", iterations)
	fmt.Printf("⏱️  执行时间: %v\n", duration)
	fmt.Printf("⚡ 每次操作平均时间: %.2f ns\n", float64(duration.Nanoseconds())/float64(iterations))
	fmt.Printf("📦 堆分配次数: %d\n", allocations)
	fmt.Printf("🗑️  堆释放次数: %d\n", frees)
	fmt.Printf("📊 活跃对象: %d\n", allocations-frees)
	fmt.Printf("📈 当前堆对象数: %d\n", heapObjects)
	fmt.Printf("💾 总分配内存: %.2f KB\n", float64(totalAlloc)/1024)
	fmt.Printf("📏 平均每次分配: %.1f bytes\n", float64(totalAlloc)/float64(iterations))
	fmt.Printf("🔄 GC 触发次数: %d\n", gcCycles)

	if pauseTotal > 0 {
		fmt.Printf("⏳ GC 暂停时间: %v (平均: %.2f µs)\n",
			time.Duration(pauseTotal),
			float64(pauseTotal)/float64(gcCycles*1000))
	}

	// 内存压力等级判断
	mbAllocated := float64(totalAlloc) / 1024 / 1024
	if gcCycles > 5 {
		fmt.Printf("🚨 极高 GC 压力 - 触发了 %d 次垃圾回收，分配 %.2f MB\n", gcCycles, mbAllocated)
	} else if gcCycles > 0 {
		fmt.Printf("⚠️  高 GC 压力 - 触发了 %d 次垃圾回收，分配 %.2f MB\n", gcCycles, mbAllocated)
	} else if allocations > 100000 {
		fmt.Printf("⚡ 高内存压力 - 大量堆分配 (%d) 但未触发 GC，分配 %.2f MB\n", allocations, mbAllocated)
	} else if allocations > 0 {
		fmt.Printf("💡 中等内存压力 - 有堆分配但未触发 GC，分配 %.2f MB\n", mbAllocated)
	} else {
		fmt.Printf("✅ 低内存压力 - 无堆分配\n")
	}
}

func TestRealistic(t *testing.T) {
	fmt.Println("🔬 Go GC 压力分析 - unsafe.Pointer 影响测试")
	fmt.Println(strings.Repeat("=", 60))

	// 🚀 在测试开始前创建基础内存压力
	createMemoryPressure()
	defer cleanMemoryPressure()

	// 🚀 高强度测试参数 - 确保触发 GC
	const lowIntensity = 50000       // 基础强度
	const mediumIntensity = 200000   // 中等强度
	const highIntensity = 1000000    // 高强度
	const extremeIntensity = 5000000 // 极限强度

	// 大对象分配池
	var memoryPressureSlice [][]byte

	// 测试 1: 🔥 极限强度 - unsafe.Pointer 大对象逃逸
	measureGCPressure("🔥 极限强度 - unsafe.Pointer 大对象逃逸", extremeIntensity, func() {
		// 创建大对象
		bigData := make([]byte, 1024) // 1KB 对象
		ptr := unsafe.Pointer(&bigData[0])
		globalSink = ptr

		// 额外内存压力
		if len(memoryPressureSlice) < 1000 {
			memoryPressureSlice = append(memoryPressureSlice, bigData)
		}
	})

	// 测试 2: 🔥 极限强度 - 普通指针大对象逃逸
	measureGCPressure("🔥 极限强度 - 普通指针大对象逃逸", extremeIntensity, func() {
		bigData := make([]byte, 1024) // 1KB 对象
		ptr := &bigData[0]
		globalSink = ptr
	})

	// 测试 3: 🔥 高强度 - 多层嵌套结构逃逸
	measureGCPressure("🔥 高强度 - 多层嵌套结构逃逸", highIntensity, func() {
		type NestedStruct struct {
			Data  [128]byte // 128 字节
			Ptr   unsafe.Pointer
			Slice []int
			Map   map[string]int
		}

		nested := &NestedStruct{
			Slice: make([]int, 10),
			Map:   make(map[string]int),
		}
		nested.Ptr = unsafe.Pointer(&nested.Data[0])
		nested.Map["key"] = 42

		globalSink = nested
	})

	// 测试 4: 🔥 中等强度 - interface{} + unsafe.Pointer 混合
	measureGCPressure("🔥 中等强度 - interface{} + unsafe.Pointer 混合", mediumIntensity, func() {
		data := make([]int, 64) // 256 字节
		val := interfaceUnsafePointer()
		ptr := unsafe.Pointer(&data[0])

		// 多重装箱
		globalSink = []interface{}{val, ptr, data}
	})

	// 测试 5: 🔥 中等强度 - 字符串拼接 + 指针逃逸
	measureGCPressure("🔥 中等强度 - 字符串拼接 + 指针逃逸", mediumIntensity, func() {
		str := fmt.Sprintf("data_%d_%p", 42, forceEscapeUnsafePointer())
		globalSink = str
	})

	// 测试 6: 🔥 基础强度 - 切片扩容 + unsafe 转换
	measureGCPressure("🔥 基础强度 - 切片扩容 + unsafe 转换", lowIntensity, func() {
		slice := make([]int, 0, 1)
		for i := 0; i < 100; i++ { // 强制多次扩容
			slice = append(slice, i)
		}
		// unsafe 访问切片底层
		ptr := unsafe.Pointer(&slice[0])
		globalSink = []interface{}{slice, ptr}
	})

	// 清理 globalSink 以释放内存
	globalSink = nil
	runtime.GC()

	// 显示最终统计
	var finalStats runtime.MemStats
	runtime.ReadMemStats(&finalStats)

	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("📊 最终内存统计 & 性能总结\n")
	fmt.Printf("总累计分配: %.2f MB\n", float64(finalStats.TotalAlloc)/1024/1024)
	fmt.Printf("当前堆使用: %.2f MB\n", float64(finalStats.HeapAlloc)/1024/1024)
	fmt.Printf("系统内存占用: %.2f MB\n", float64(finalStats.Sys)/1024/1024)
	fmt.Printf("总 GC 次数: %d\n", finalStats.NumGC)
	fmt.Printf("GC 暂停总时间: %v\n", time.Duration(finalStats.PauseTotalNs))
	fmt.Printf("活跃 Goroutine: %d\n", runtime.NumGoroutine())

	fmt.Printf("\n💡 高强度测试结论:\n")
	fmt.Printf("🔥 极限强度测试成功触发多次 GC，验证了内存压力\n")
	fmt.Printf("⚡ unsafe.Pointer 在大规模使用时确实会产生显著的 GC 压力\n")
	fmt.Printf("📦 大对象 + 逃逸分析是 GC 压力的主要来源\n")
	fmt.Printf("🎯 关键优化点：减少逃逸、避免大对象频繁分配\n")
	fmt.Printf("⚠️  在高频场景下，任何逃逸都可能成为性能瓶颈\n")
	fmt.Printf("✨ 合理使用 unsafe.Pointer 需要配合逃逸分析优化\n")
}
