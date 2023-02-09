package stack

import "errors"

// Stack a first in last out data structure, based on slice
type Stack[T any] struct {
	array    []T
	stackTop int
}

// NewStack return a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		array:    make([]T, 0),
		stackTop: 0,
	}
}

// Push put a new value into the stack
func (s *Stack[T]) Push(v T) {
	if s.stackTop >= len(s.array) {
		s.array = append(s.array, v) // use append to expand the underlying slice
	} else {
		s.array[s.stackTop-1] = v
	}
	s.stackTop++
}

// Pop remove the top of the stack
func (s *Stack[T]) Pop() error {
	if s.stackTop == 0 {
		return errors.New("the stack is empty")
	}
	s.stackTop--
	return nil
}

// Top return the top of the stack
func (s *Stack[T]) Top() (T, error) {
	var zeroT T
	if s.stackTop == 0 {
		return zeroT, errors.New("the stack is empty")
	}
	return s.array[s.stackTop-1], nil
}

// Size return the size of the stack
func (s *Stack[T]) Size() int {
	return s.stackTop
}

// Empty return the stack is empty or not
func (s *Stack[T]) Empty() bool {
	return s.stackTop == 0
}
