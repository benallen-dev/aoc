package stack

import (
	"errors"
	"fmt"
)

// It occurs to me that instead of using a slice you could use a linked list
// for O(1) pushes and pops, because you likely don't care about the rest of
// the state anyway - but I'm too lazy to change that until all the memory
// allocation stuff starts to become an issue.

type Stack[T any] struct{
	items []T
	size int
	Formatter func(v T) string
}

// I should try the WithFoo pattern for changing the formatter instead of just exposing internals
func New [T any](items []T) Stack[T] {
	return Stack[T] {
		items: items,
		size: len(items),
	}
}

func (s *Stack[T]) String() string {
	out := ""
	
	for i, v := range s.items {
		if i == 0 {
			out += "["
		} else if i < len(s.items) {
			out += ", "
		}

		if s.Formatter != nil {
			out += s.Formatter(v)
		} else {
			out += fmt.Sprintf("%v", v)
		}

		if i == len(s.items) -1 {
			out += "]"
		}
	}

	return out
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
	s.size = s.size + 1
}

func (s *Stack[T]) Pop() (T, error) {
	if s.isEmpty() {
		var zero T
		return zero, errors.New("attempted to pop from an empty stack")
	}

	p := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	s.size = s.size -1

	return p, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.isEmpty() {
		var zero T
		return zero, errors.New("attempted to pop from an empty stack")
	}

	return s.items[s.size-1], nil
}

func (s *Stack[T]) isEmpty() bool {
	return s.size == 0
}

