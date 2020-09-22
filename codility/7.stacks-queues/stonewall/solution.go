package stonewall

func solution(H []int) int {

	l := NewIntStack(len(H))
	result := 0

	for _, value := range H {
		for l.size > 0 && l.Front() > value {
			l.Pop()
		}

		if l.size == 0 || l.Front() < value {
			l.Push(value)
			result++
		}
	}

	return result
}
