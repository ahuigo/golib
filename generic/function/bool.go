package function

func And(elem1 bool) func(bool) bool {
	return func(elem0 bool) bool {
		return elem0 && elem1
	}
}

func Or(elem1 bool) func(bool) bool {
	return func(elem0 bool) bool {
		return elem0 || elem1
	}
}

func Not(elem0 bool) bool {
	return !elem0
}
