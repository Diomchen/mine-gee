package set

type Set[T comparable] struct {
	element map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		element: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(e T) {
	s.element[e] = struct{}{}
}

func (s *Set[T]) Contains(e T) bool {
	_, ok := s.element[e]
	return ok
}

func (s *Set[T]) Remove(e T) {
	delete(s.element, e)
}

func (s *Set[T]) Size() int {
	return len(s.element)
}

func (s *Set[T]) Print() {
	for e := range s.element {
		print(e, " ")
	}
}
