package function

func Length(array string) int {
	return len(array)
}

func LengthA[T any](array []T) int {
	return len(array)
}

func LengthM[T comparable, U any](m map[T]U) int {
	return len(m)
}
