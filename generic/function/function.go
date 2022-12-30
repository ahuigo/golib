package function

func Absurd[T any](_ T) {
}

func Identity[T any](arg T) T {
	return arg
}

func Const[T0, T1 any](arg0 T0) func(T1) T0 {
	return func(_ T1) T0 {
		return arg0
	}
}

func First[T0, T1 any](arg0 T0, _ T1) T0 {
	return arg0
}

func Second[T0, T1 any](_ T0, arg1 T1) T1 {
	return arg1
}

func Compose[T0, T1, T2 any](f0 func(T1) T2, f1 func(T0) T1) func(T0) T2 {
	return func(t0 T0) T2 {
		return f0(f1(t0))
	}
}

func Then[T0, T1, T2 any](f0 func(T0) T1, f1 func(T1) T2) func(T0) T2 {
	return Compose(f1, f0)
}

func Curry2[T0, T1, R any](f func(T0, T1) R) func(T0) func(T1) R {
	return func(t0 T0) func(T1) R {
		return func(t1 T1) R {
			return f(t0, t1)
		}
	}
}

func UnCurry2[T0, T1, R any](f func(T0) func(T1) R) func(T0, T1) R {
	return func(t0 T0, t1 T1) R {
		return f(t0)(t1)
	}
}

func Flip[T0, T1, R any](f func(T0, T1) R) func(T1, T0) R {
	return func(t1 T1, t0 T0) R {
		return f(t0, t1)
	}
}

func Fix[T comparable](f func(T) T) func(T) T {
	return func(arg T) T {
		var cache = arg
		for v := f(arg); v != cache; v = f(v) {
			cache = v
		}
		return cache
	}
}

func On[T0, T1, T2 any](f func(T1, T1) T2, argF func(T0) T1) func(T0, T0) T2 {
	return func(arg0, arg1 T0) T2 {
		return f(argF(arg0), argF(arg1))
	}
}

func Call[T0, T1 any](f func(T0) T1, arg0 T0) T1 {
	return f(arg0)
}
