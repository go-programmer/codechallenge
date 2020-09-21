package brackets

func thirdpartySolution(S string) int {
	l := NewStringStack(len(S))
	pairs := make(map[string]string)
	pairs["{"] = "}"
	pairs["["] = "]"
	pairs["("] = ")"

	if len(S) == 0 {
		return 1
	}

	if len(S)%2 != 0 {
		return 0
	}

	for _, value := range S {
		val := string(value)
		if val == "(" || val == "{" || val == "[" {
			l.Push(val)
		} else if val == ")" || val == "}" || val == "]" {
			if l.size == 0 {
				return 0
			}

			head := l.Front()
			if pairs[head] == val {
				l.Pop()
			} else {
				return 0
			}
		}
	}

	if l.size == 0 {
		return 1
	} else {
		return 0
	}
}

// StringStack data structure
type StringStack struct {
	size int
	data []string
}

// NewStringStack initialize StringStack
func NewStringStack(len int) *StringStack {
	return &StringStack{
		size: 0,
		data: make([]string, len),
	}
}

// Push string into StringStack
func (s *StringStack) Push(item string) {
	if s.size < len(s.data) {
		s.data[s.size] = item
		s.size++
	}
}

// Pop from StringStack
func (s *StringStack) Pop() string {
	if s.size > 0 {
		item := s.data[s.size-1]
		s.size--

		return item
	} else {
		return ""
	}
}

// Front returns from top of stack
func (s *StringStack) Front() string {
	if s.size > 0 {
		return s.data[s.size-1]
	} else {
		return ""
	}
}
