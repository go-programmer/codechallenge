package nesting

func solution(S string) int {

	lenS := len(S)

	if lenS == 0 {
		return 1
	}

	ss := NewStack(lenS / 2)

	for _, v := range S {

		if v == 40 {

			if err := ss.Push(int8(v)); err != nil {
				return 0
			}

		} else if v == 41 {

			if _, err := ss.Pop(); err != nil {
				return 0
			}

		}
	}

	if ss.size <= 0 {
		return 1
	}

	return 0
}
