package function

import "testing"

/*
refer: https://www.reddit.com/r/golang/comments/tyn2c9/questionable_generics_cannot_use_type_aliases/
*/
type A struct {
	Value int
}

type ASlice []A

func PrintSlice[T any](s []T) {

}

// func PrintSlicePtrOld[T any](s *[]T) {} // error
func PrintSlicePtr[T any, S ~[]T](s *S) {}

func TestPtrSlice(t *testing.T) {
	slice := ASlice{A{Value: 1}, A{Value: 2}, A{Value: 3}}
	// works,
	PrintSlice(slice) // slice 是named ASlice->unamed []A, underlying: []A -> []A

	// 😤 compile error: cannot use &slice (value of type *ASlice) as type *[]A in argument to PrintSlice
	// PrintSlicePtrOld(&slice) //&slice是 unamed *ASlice->unamed *[]A, underlying: *ASlice -> *[]A
	PrintSlicePtr(&slice) //&slice是 unamed *ASlice->unamed *[]A, underlying: *ASlice -> *[]A (但是有 ~[]T, 所以可以匹配)

	sameSliceButDifferent := []A{{Value: 1}, {Value: 2}, {Value: 3}}
	// works just fine
	PrintSlicePtr(&sameSliceButDifferent) //  unamed *[]A->unamed *[]A, underlying: *[]A-> *[]A

}
