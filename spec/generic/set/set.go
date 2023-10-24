package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return map[T]struct{}{}
}

func (s Set[T]) Add(elem T) Set[T] {
	s[elem] = struct{}{}
	return s
}

func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}
