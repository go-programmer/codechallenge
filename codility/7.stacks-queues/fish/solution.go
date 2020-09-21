package fish

// Check if fish is downsteam and add to stack.
// If fish is upstream, check how many it will eat from downstream.
func solution(A []int, B []int) int {
	stack := NewIntStack(len(A))
	response := 0

	for i := range B {

		if B[i] == 1 {
			stack.Push(A[i])

		} else {

			// Kill all sharks moving up.
			for stack.size > 0 {

				if stack.Front() < A[i] {
					stack.Pop()
				} else {
					break
				}
			}

			// Upstream fish kills all downstream fish
			// so there is one fish that lives.
			if stack.size == 0 {
				response++
			}
		}
	}

	return response + stack.size
}

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
