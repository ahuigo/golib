package collections

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	fn "demo/function"
	"demo/maths"
)

func TestInclude(t *testing.T) {
	assert.True(t, Include([]string{"a", "b", "c"}, "b"))
	assert.False(t, Include([]string{"a", "b", "c"}, "d"))
}

func TestMap(t *testing.T) {
	assert.Equal(t, Map([]string{"a", "ab", "abc"}, fn.Length), []int{1, 2, 3})
	assert.Equal(t, Map([]int{1, 2, 3}, strconv.Itoa), []string{"1", "2", "3"})
}

func TestFilter(t *testing.T) {
	assert.Equal(t, Filter([]string{"a", "ab", "abc"},
		fn.Then(fn.Length, fn.GreaterThan(1))), []string{"ab", "abc"})
	assert.Equal(t, Filter([]string{"a", "ab", "abc"},
		fn.Then(fn.Length, fn.LessEqualThan(2))), []string{"a", "ab"})
	assert.Equal(t, Filter([]string{"a", "ab", "abc"},
		fn.Then(fn.Then(fn.Length, fn.Equal(2)), fn.Not)), []string{"a", "abc"})
}

func TestZipWith(t *testing.T) {
	assert.Equal(t, ZipWith(
		[]int{1, 2, 3},
		[]string{"a", "ab"},
		func(t0 int, t1 string) int {
			return t0 + len(t1)
		},
	), []int{2, 4})
}

func TestGroupBy(t *testing.T) {
	assert.Equal(t, GroupBy(
		[]string{"a", "ab", "c", "ac", "def"},
		fn.Length,
	), map[int][]string{
		1: {"a", "c"},
		2: {"ab", "ac"},
		3: {"def"},
	})
}

func TestDistinct(t *testing.T) {
	assert.Equal(t, Distinct([]int{1, 2, 1, 3, 4}), []int{1, 2, 3, 4})
}

func TestFoldLeft(t *testing.T) {
	assert.Equal(t, FoldLeft([]int{3, 1, 2}, 0, maths.Max[int]), 3)
	assert.Equal(t, FoldLeft([]int{3, 1, 2}, 4, maths.Min[int]), 1)
	assert.Equal(t, FoldLeft([]string{"a", "b", "c"}, "d", fn.StrCat), "dabc")
}

func TestReduceLeft(t *testing.T) {
	assert.Equal(t, ReduceLeft([]int{3, 1, 2}, maths.Max[int]), 3)
	assert.Equal(t, ReduceLeft([]int{3, 1, 2}, maths.Min[int]), 1)
	assert.Equal(t, ReduceLeft([]string{"a", "b", "c"}, fn.StrCat), "abc")
}

func TestFoldRight(t *testing.T) {
	assert.Equal(t, FoldRight([]int{3, 1, 2}, 0, maths.Max[int]), 3)
	assert.Equal(t, FoldRight([]int{3, 1, 2}, 4, maths.Min[int]), 1)
	assert.Equal(t, FoldRight([]string{"a", "b", "c"}, "d", fn.StrCat), "abcd")
}

func TestReduceRight(t *testing.T) {
	assert.Equal(t, ReduceRight([]int{3, 1, 2}, maths.Max[int]), 3)
	assert.Equal(t, ReduceRight([]int{3, 1, 2}, maths.Min[int]), 1)
	assert.Equal(t, ReduceRight([]string{"a", "b", "c"}, fn.StrCat), "abc")
}

func TestReplicate(t *testing.T) {
	assert.Equal(t, Replicate(0, 1), []int{})
	assert.Equal(t, Replicate(-1, 1), []int{})
	assert.Equal(t, Replicate(3, 1), []int{1, 1, 1})
}

func TestScanLeft(t *testing.T) {
	assert.Equal(t, ScanLeft([]int{1, 2, 3}, 1, fn.UnCurry2(fn.Add[int])), []int{1, 2, 4, 7})
	assert.Equal(t, ScanLeft([]int{1, 1, 1}, 0, fn.UnCurry2(fn.Add[int])), []int{0, 1, 2, 3})
}

func TestScanLeft1(t *testing.T) {
	assert.Equal(t, ScanLeft1([]int{1, 2, 3}, fn.UnCurry2(fn.Add[int])), []int{1, 3, 6})
	assert.Equal(t, ScanLeft1([]int{1, 1, 1}, fn.UnCurry2(fn.Add[int])), []int{1, 2, 3})
}

func TestScanRight(t *testing.T) {
	assert.Equal(t, ScanRight([]int{1, 2, 3}, 1, fn.UnCurry2(fn.Add[int])), []int{7, 6, 4, 1})
	assert.Equal(t, ScanRight([]int{1, 1, 1}, 0, fn.UnCurry2(fn.Add[int])), []int{3, 2, 1, 0})
	assert.Equal(t, ScanRight([]int{}, 0, fn.UnCurry2(fn.Add[int])), []int{0})
}

func TestScanRight1(t *testing.T) {
	assert.Equal(t, ScanRight1([]int{1, 2, 3}, fn.UnCurry2(fn.Add[int])), []int{6, 5, 3})
	assert.Equal(t, ScanRight1([]int{1, 1, 1}, fn.UnCurry2(fn.Add[int])), []int{3, 2, 1})
	assert.Equal(t, ScanRight1([]int{0}, fn.UnCurry2(fn.Add[int])), []int{0})
}

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse([]int{1, 2, 3, 4}), []int{4, 3, 2, 1})
	assert.Equal(t, Reverse([]int{1}), []int{1})
	assert.Equal(t, Reverse([]int{}), []int{})
	assert.Equal(t, Reverse([]string{"a", "b", "c"}), []string{"c", "b", "a"})
}

func TestTakeWhile(t *testing.T) {
	assert.Equal(t,
		TakeWhile(
			[]string{"a", "ab", "abc"},
			fn.UnCurry2(
				fn.Const[func(string) bool, int](fn.Then(fn.Length, fn.LessThan(3))),
			),
		),
		[]string{"a", "ab"},
	)
}

func TestTake(t *testing.T) {
	assert.Equal(t, Take(3, []int{1, 2, 3, 4, 5}), []int{1, 2, 3})
	assert.Equal(t, Take(0, []int{1, 2, 3, 4, 5}), []int{})
	assert.Equal(t, Take(-1, []int{1, 2, 3, 4, 5}), []int{})
	assert.Equal(t, Take(10, []int{1, 2, 3, 4, 5}), []int{1, 2, 3, 4, 5})
	assert.Equal(t, Take(2, []int{1, 2, 3}), []int{1, 2})
	assert.Equal(t, Take(3, []int{1, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, Take(4, []int{1, 2, 3}), []int{1, 2, 3})
}

func TestDropWhile(t *testing.T) {
	assert.Equal(t,
		DropWhile(
			[]string{"a", "b", "ab", "ac"},
			fn.UnCurry2(
				fn.Const[func(string) bool, int](fn.Then(fn.Length, fn.LessThan(2)))),
		),
		[]string{"ab", "ac"},
	)
}

func TestDrop(t *testing.T) {
	assert.Equal(t, Drop(0, []int{1, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, Drop(-1, []int{1, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, Drop(1, []int{1, 2, 3}), []int{2, 3})
	assert.Equal(t, Drop(3, []int{1, 2, 3}), []int{})
	assert.Equal(t, Drop(4, []int{1, 2, 3}), []int{})
}

func TestAll(t *testing.T) {
	assert.False(t, All([]int{1, 2, 3}, fn.Then(fn.Mod(2), fn.Equal(0))))
	assert.True(t, All([]int{2, 4, 6}, fn.Then(fn.Mod(2), fn.Equal(0))))
	assert.True(t, All([]int{}, fn.Then(fn.Mod(2), fn.Equal(0))))
}

func TestAny(t *testing.T) {
	assert.True(t, Any([]int{1, 2, 3}, fn.Then(fn.Mod(2), fn.Equal(0))))
	assert.False(t, Any([]int{1, 3, 5}, fn.Then(fn.Mod(2), fn.Equal(0))))
	assert.False(t, Any([]int{}, fn.Then(fn.Mod(2), fn.Equal(0))))
}

func TestMaximum(t *testing.T) {
	assert.Equal(t, Maximum([]int{1, 3, 2}), 3)
	assert.Equal(t, Maximum([]int{1, 2, 3}), 3)
	assert.Equal(t, Maximum([]int{1}), 1)
	assert.Equal(t, Maximum([]int{1, -1, 1, -2}), 1)
}

func TestMinimum(t *testing.T) {
	assert.Equal(t, Minimum([]int{3, 1, 2}), 1)
	assert.Equal(t, Minimum([]int{3, 2, 1}), 1)
	assert.Equal(t, Minimum([]int{1}), 1)
	assert.Equal(t, Minimum([]int{1, -1, 1, -2}), -2)
}

func TestInit(t *testing.T) {
	assert.Equal(t, Init([]string{"a", "b", "c"}), []string{"a", "b"})
	assert.Equal(t, Init([]string{"a"}), []string{})
	assert.Equal(t, Init([]string{}), []string{})
}

func TestTail(t *testing.T) {
	assert.Equal(t, Tail([]string{"a", "b", "c"}), []string{"b", "c"})
	assert.Equal(t, Tail([]string{"a", "b"}), []string{"b"})
	assert.Equal(t, Tail([]string{"a"}), []string{})
	assert.Equal(t, Tail([]string{}), []string{})
}

func TestHead(t *testing.T) {
	assert.Equal(t, Head([]string{"a", "b", "c"}), "a")
	assert.Equal(t, Head([]string{"a", "b"}), "a")
}

func TestLast(t *testing.T) {
	assert.Equal(t, Last([]string{"a", "b", "c"}), "c")
	assert.Equal(t, Last([]string{"a", "b"}), "b")
}

func TestSpan(t *testing.T) {
	odds, evens := Span(
		[]int{1, 2, 3, 4, 5},
		fn.UnCurry2(fn.Const[func(int) bool, int](fn.Then(fn.Mod(2), fn.Equal(0)))),
	)
	assert.Equal(t, []int{1, 3, 5}, odds)
	assert.Equal(t, []int{2, 4}, evens)
}

func TestSplitAt(t *testing.T) {
	l, r := SplitAt([]int{0, 1, 2, 3, 4}, 2)
	assert.Equal(t, []int{0, 1}, l)
	assert.Equal(t, []int{2, 3, 4}, r)
}

func TestSum(t *testing.T) {
	assert.Equal(t, Sum([]float64{1, 2, 3}), float64(6))
	assert.Equal(t, Sum([]float64{}), float64(0))
}

func TestForEach(t *testing.T) {
	var i = 0
	ForEach([]int{1, 2, 3}, func(_, e int) {
		i += e
	})
	assert.Equal(t, 6, i)
}

func TestAnd(t *testing.T) {
	assert.Equal(t, And(true, false), false)
	assert.Equal(t, And(true, true), true)
}

func TestOr(t *testing.T) {
	assert.Equal(t, Or(true, false), true)
	assert.Equal(t, Or(false, false), false)
	assert.Equal(t, Or(true, true), true)
}

func TestConcat(t *testing.T) {
	assert.Equal(t, Concat([]int{1}, []int{2, 3}, []int{4}), []int{1, 2, 3, 4})
	assert.Equal(t, Concat([]int{1}, []int{}, []int{4}), []int{1, 4})
	assert.Equal(t, Concat([]int{1}, nil, []int{4}), []int{1, 4})
}

func TestConcatMap(t *testing.T) {
	assert.Equal(t,
		ConcatMap[int, string]([]int{1, 2, 3}, fn.Curry2(fn.Flip(Replicate[string]))("a")),
		[]string{"a", "a", "a", "a", "a", "a"},
	)
}

func TestSort(t *testing.T) {
	assert.Equal(t, Sort([]int{2, 1, 4, 3}), []int{1, 2, 3, 4})
}
