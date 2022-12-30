package function

import "golang.org/x/exp/constraints"

func Equal[T comparable](elem1 T) func(T) bool {
	return func(elem0 T) bool {
		return elem0 == elem1
	}
}

func GreaterThan[T constraints.Ordered](elem1 T) func(T) bool {
	return func(elem0 T) bool {
		return elem0 > elem1
	}
}

func LessThan[T constraints.Ordered](elem1 T) func(T) bool {
	return func(elem0 T) bool {
		return elem0 < elem1
	}
}

func GreaterEqualThan[T constraints.Ordered](elem1 T) func(T) bool {
	return func(elem0 T) bool {
		return elem0 >= elem1
	}
}

func LessEqualThan[T constraints.Ordered](elem1 T) func(T) bool {
	return func(elem0 T) bool {
		return elem0 <= elem1
	}
}
