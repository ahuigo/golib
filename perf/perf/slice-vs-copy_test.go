// refer: https://geektutu.com/post/hpg-slice.html
package perf

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func lastNumsBySlice(origin []int) []int {
	// 切片引用数组在内存中仍然引用了大量空间，得不到释放。比较推荐的做法，使用 copy 替代 re-slice。
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}

// 8*128*1024=1M
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func printMem(t *testing.T, args ...string) {
	t.Helper()
	var rtm runtime.MemStats
	// 打印运行时内存大小
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB(%v)", float64(rtm.Alloc)/1024./1024., args)
}

func testLastChars(t *testing.T, f func([]int) []int) {
	runtime.GC()
	printMem(t, "start")
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		//printMem(t, fmt.Sprintf("start %d", k))
		origin := generateWithCap(128 * 1024) // 1M
		ans = append(ans, f(origin))
		//runtime.GC() // 1. 循环内gc不能回收origin(切片ans还有引用)
	}
	needGc := false
	if needGc {
		runtime.GC() // 2. 显式gc：循环外gc可以回收origin(主动gc, 切片无引用)
	} else {
		println("不触发GC") // 3. 触发隐式GC: 默认触发条件是内存增长50% 且进入函数调用时触发
	}
	printMem(t, "end")
	_ = ans
}

func TestLastCharsBySlice(t *testing.T) { testLastChars(t, lastNumsBySlice) }
func TestLastCharsByCopy(t *testing.T)  { testLastChars(t, lastNumsByCopy) }
