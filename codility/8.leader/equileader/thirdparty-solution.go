package equileader

import "math"

// We utilize the Leader implementation internals
// to maintain a list of subleaders in a map
func thirdpartySolution(A []int) int {
	leadersCount := 0
	arrayLen := len(A)
	l := NewIntStack(arrayLen)
	candidate := -1
	leader := -1
	count := 0
	leftLeadersCount := 0
	leftSequenceLength := 0
	rightSequenceLength := 0

	if arrayLen == 1 {
		return 0
	}

	for _, value := range A {
		//O(n)
		if l.size == 0 {
			l.Push(value)
		} else {
			if value != l.Front() {
				l.Pop()
			} else {
				l.Push(value)
			}
		}
	}

	if l.size > 0 {
		candidate = l.Front()
	} else {
		return 0
	}

	for _, value := range A {
		if value == candidate {
			count += 1
		}
	}

	if count > int(math.Floor(float64(arrayLen)/2.0)) {
		leader = candidate
	} else {
		return 0
	}

	for i, value := range A {
		if value == leader {
			// The key here is to compare the current leader with the current value
			leftLeadersCount += 1
		}

		/* validate left sequence */
		leftSequenceLength = i + 1

		if leftLeadersCount > int(math.Floor(float64(leftSequenceLength)/2.0)) {
			/* validate right sequence */
			rightSequenceLength = arrayLen - leftSequenceLength

			// Here we check if the remaining leaders count on the left side is more than have of the remaining on the right side
			if count-leftLeadersCount > int(math.Floor(float64(rightSequenceLength)/2.0)) {
				/* both sequences have valid leaders of the same value */
				leadersCount += 1
			}
		}
	}

	return leadersCount
}

func Leader(A []int) int {
	// O(N)
	arrayLen := len(A)

	if arrayLen == 1 {
		return A[0]
	}

	l := NewIntStack(arrayLen)
	candidate := -1
	leader := -1
	count := 0

	for _, value := range A {
		//O(n)
		if l.size == 0 {
			l.Push(value)
		} else {
			if value != l.Front() {
				l.Pop()
			} else {
				l.Push(value)
			}
		}
	}

	if l.size > 0 {
		candidate = l.Front()
	}

	for _, value := range A {
		if value == candidate {
			count += 1
		}
	}

	if count > int(math.Floor(float64(arrayLen)/2.0)) {
		leader = candidate
	}

	return leader

}

type IntStack struct {
	size int
	data []int
}

func NewIntStack(len int) *IntStack {
	return &IntStack{
		size: 0,
		data: make([]int, len),
	}
}

func (s *IntStack) Push(item int) {
	if s.size < len(s.data) {
		s.data[s.size] = item
		s.size++
	}
}

func (s *IntStack) Pop() int {
	item := s.data[s.size-1]
	s.size -= 1

	return item
}

func (s *IntStack) Front() int {
	return s.data[s.size-1]
}
