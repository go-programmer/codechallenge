package nesting

import "errors"

// stack data structure
type stack struct {
	size int
	data []int8
}

// NewStack initialize StringStack
func NewStack(len int) *stack {
	return &stack{
		size: -1,
		data: make([]int8, len),
	}
}

// Push string into stack
func (s *stack) Push(item int8) error {
	s.size++

	if s.size >= len(s.data) {
		return errors.New("Stack size exceeded")
	}

	s.data[s.size] = item
	return nil
}

// Pop from stack
func (s *stack) Pop() (int8, error) {

	isEmpty, err := stackEmpty(s)
	if isEmpty {
		return -1, err
	}
	poppedVal := s.data[s.size]
	s.data[s.size] = 0
	s.size--

	return poppedVal, nil
}

// Front returns from top of stack
func (s *stack) Front() (int8, error) {

	isEmpty, err := stackEmpty(s)
	if isEmpty {
		return -1, err
	}

	return s.data[s.size-1], nil

}

// stackEmpty Check if stack is empty.
func stackEmpty(s *stack) (bool, error) {

	if s.size < 0 {
		return true, errors.New("Stack is empty")
	}

	return false, nil
}
