package maths

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a T, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

func Floor[T constraints.Integer](v float64) T {
	return T(math.Floor(v))
}

func Ceil[T constraints.Integer](v float64) T {
	return T(math.Ceil(v))
}
