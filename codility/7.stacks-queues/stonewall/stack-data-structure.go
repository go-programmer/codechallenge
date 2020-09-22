package stonewall

// Integer Stack Data Structure

// IntStack int stack
type IntStack struct {
	size int
	data []int
}

// NewIntStack int stack constructor
func NewIntStack(len int) *IntStack {
	return &IntStack{
		size: 0,
		data: make([]int, len),
	}
}

// Push item to stack
func (s *IntStack) Push(item int) {
	if s.size < len(s.data) {
		s.data[s.size] = item
		s.size++
	}
}

// Pop item from stack
func (s *IntStack) Pop() int {
	item := s.data[s.size-1]
	s.size--

	return item
}

// Front Get the front of stack
func (s *IntStack) Front() int {
	return s.data[s.size-1]
}
