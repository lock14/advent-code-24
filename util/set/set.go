package set

import "iter"

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(es ...T) {
	for _, e := range es {
		s.m[e] = struct{}{}
	}

}

func (s *Set[T]) Remove(e T) {
	delete(s.m, e)
}

func (s *Set[T]) Contains(e T) bool {
	_, ok := s.m[e]
	return ok
}

func (s *Set[T]) Empty() bool {
	return len(s.m) == 0
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range s.m {
			if !yield(e) {
				return
			}
		}
	}
}
