package set

type Set[T comparable] map[T]bool


func (s Set[T]) Add(v T)  {
	s[v] = true
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Length() int {
	return len(s)
}

