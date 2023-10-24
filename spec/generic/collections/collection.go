package collections

import (
	"sort"

	"golang.org/x/exp/constraints"

	fn "demo/function"

	"demo/maths"
)

func Include[T comparable](array []T, element T) bool {
	for _, elem := range array {
		if elem == element {
			return true
		}
	}
	return false
}

func Map[T any, U any](array []T, op func(T) U) []U {
	rs := make([]U, len(array))
	for i, elem := range array {
		rs[i] = op(elem)
	}
	return rs
}

func Filter[T any](array []T, pred func(T) bool) []T {
	rs := make([]T, 0)
	for _, elem := range array {
		_elem := elem
		if pred(_elem) {
			rs = append(rs, _elem)
		}
	}
	return rs
}

func ZipWith[T0 any, T1 any, R any](array0 []T0, array1 []T1, zip func(T0, T1) R) []R {
	rs := make([]R, 0)
	_len := maths.Min[int](len(array0), len(array1))
	for i := 0; i < _len; i++ {
		rs = append(rs, zip(array0[i], array1[i]))
	}
	return rs
}

func GroupBy[T any, U comparable](array []T, keyMapper func(T) U) map[U][]T {
	rm := make(map[U][]T)
	for _, elem := range array {
		_elem := elem
		key := keyMapper(_elem)
		rm[key] = append(rm[key], _elem)
	}
	return rm
}

func Distinct[T comparable](array []T) []T {
	cache := make(map[T]struct{})
	rs := make([]T, 0)
	for _, elem := range array {
		_elem := elem
		if _, ok := cache[_elem]; ok {
			continue
		} else {
			rs = append(rs, _elem)
			cache[_elem] = struct{}{}
		}
	}
	return rs
}

func FoldLeft[R any, T any](array []T, first R, op func(R, T) R) R {
	r := first
	for _, elem := range array {
		_elem := elem
		r = op(r, _elem)
	}
	return r
}

func ReduceLeft[T any](array []T, op func(T, T) T) T {
	return FoldLeft(array[1:], array[0], op)
}

func FoldRight[T any, R any](array []T, last R, op func(T, R) R) R {
	r := last
	_len := len(array)
	for i := _len - 1; i >= 0; i-- {
		r = op(array[i], r)
	}
	return r
}

func ReduceRight[T any](array []T, op func(T, T) T) T {
	_len := len(array)
	return FoldRight(array[:_len-1], array[_len-1], op)
}

func Replicate[T any](n int, elem T) []T {
	var rs = make([]T, 0)
	if n <= 0 {
		return rs
	}
	for i := 0; i < n; i++ {
		rs = append(rs, elem)
	}
	return rs
}

func ScanLeft[T, R any](array []T, first R, op func(R, T) R) []R {
	rs := make([]R, 0, len(array)+1)
	rs = append(rs, first)
	r := first
	for i := range array {
		r = op(r, array[i])
		rs = append(rs, r)
	}
	return rs
}

func ScanLeft1[T any](array []T, op func(T, T) T) []T {
	return ScanLeft(array[1:], array[0], op)
}

func ScanRight[T, R any](array []T, last R, op func(T, R) R) []R {
	_len := len(array)
	rs := make([]R, _len+1)
	rs[_len] = last
	l := last
	for i := _len - 1; i >= 0; i-- {
		l = op(array[i], l)
		rs[i] = l
	}
	return rs
}

func ScanRight1[T any](array []T, op func(T, T) T) []T {
	return ScanRight(array[:len(array)-1], array[len(array)-1], op)
}

func Reverse[T any](array []T) []T {
	rs := make([]T, len(array))
	_len := len(rs)
	for i := range array {
		rs[_len-i-1] = array[i]
	}
	return rs
}

func TakeWhile[T any](array []T, pred func(int, T) bool) []T {
	rs := make([]T, 0)
	for i := range array {
		if pred(i, array[i]) {
			rs = append(rs, array[i])
		} else {
			return rs
		}
	}
	return rs
}

func Take[T any](n int, array []T) []T {
	rs := make([]T, 0)
	if n <= 0 {
		return rs
	}
	for i, cnt := 0, 0; i < len(array) && cnt < n; i++ {
		rs = append(rs, array[i])
		cnt++
	}
	return rs
}

func DropWhile[T any](array []T, pred func(int, T) bool) []T {
	for i := range array {
		if pred(i, array[i]) {
			continue
		} else {
			return array[i:]
		}
	}
	return make([]T, 0)
}

func Drop[T any](n int, array []T) []T {
	if n <= 0 {
		return array[:]
	}
	if n > len(array) {
		return []T{}
	}
	return array[n:]
}

func All[T any](array []T, pred func(T) bool) bool {
	for _, elem := range array {
		if !pred(elem) {
			return false
		}
	}
	return true
}

func Any[T any](array []T, pred func(T) bool) bool {
	for _, elem := range array {
		if pred(elem) {
			return true
		}
	}
	return false
}

func Minimum[T constraints.Ordered](array []T) T {
	min := array[0]
	for _, _elem := range array[1:] {
		min = maths.Min(min, _elem)
	}
	return min
}

func Maximum[T constraints.Ordered](array []T) T {
	max := array[0]
	for _, _elem := range array[1:] {
		max = maths.Max(max, _elem)
	}
	return max
}

func Init[T any](array []T) []T {
	if len(array) == 0 {
		return []T{}
	}
	return array[:len(array)-1]
}

func Tail[T any](array []T) []T {
	if len(array) == 0 {
		return []T{}
	}
	return array[1:]
}

func Head[T any](array []T) T {
	return array[0]
}

func Last[T any](array []T) T {
	return array[len(array)-1]
}

func Span[T any](array []T, pred func(int, T) bool) ([]T, []T) {
	ls := make([]T, 0)
	rs := make([]T, 0)
	for i, elem := range array {
		_elem := elem
		if pred(i, _elem) {
			rs = append(rs, _elem)
		} else {
			ls = append(ls, _elem)
		}
	}
	return ls, rs
}

func SplitAt[T any](array []T, n int) ([]T, []T) {
	return Span(array, func(index int, _ T) bool {
		return index >= n
	})
}

func ForEach[T any](array []T, op func(int, T)) {
	for i, elem := range array {
		_elem := elem
		op(i, _elem)
	}
}

func Sum[T constraints.Ordered](array []T) T {
	var zero T
	return FoldLeft(array, zero, func(e0, e1 T) T {
		return e0 + e1
	})
}

func Or(array ...bool) bool {
	return Any(array, fn.Identity[bool])
}

func And(array ...bool) bool {
	return All(array, fn.Identity[bool])
}

func Concat[T any](arrays ...[]T) []T {
	rs := make([]T, 0)
	for _, elem := range arrays {
		rs = append(rs, elem...)
	}
	return rs
}

func ConcatMap[T, U any](array []T, op func(T) []U) []U {
	return Concat(Map(array, op)...)
}

type sorter[T constraints.Ordered] []T

// var _ sort.Interface = (*sorter[])(nil)

func (s sorter[T]) Len() int {
	return len(s)
}

func (s sorter[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sorter[T]) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

func Sort[T constraints.Ordered](array []T) []T {
	sort.Sort(sorter[T](array))
	return array
}
