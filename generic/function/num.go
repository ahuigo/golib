package function

import "golang.org/x/exp/constraints"

func Add[T constraints.Ordered](elem1 T) func(T) T {
	return func(elem0 T) T {
		return elem0 + elem1
	}
}

func Sub[T constraints.Integer | constraints.Float | constraints.Complex](elem1 T) func(T) T {
	return func(elem0 T) T {
		return elem0 - elem1
	}
}

func Mul[T constraints.Integer | constraints.Float | constraints.Complex](elem1 T) func(T) T {
	return func(elem0 T) T {
		return elem0 * elem1
	}
}

func Div[T constraints.Integer | constraints.Float | constraints.Complex](elem1 T) func(T) T {
	return func(elem0 T) T {
		return elem0 / elem1
	}
}

func Mod[T constraints.Integer](elem1 T) func(T) T {
	return func(elem0 T) T {
		return elem0 % elem1
	}
}
