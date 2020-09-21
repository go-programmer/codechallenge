package brackets

func solution(S string) int {

	lenS := len(S)

	if lenS == 0 {
		return 1
	}

	// If first element is closing.
	if S[0] == 125 || S[0] == 41 || S[0] == 93 {
		return 0
	}

	// If last is an opening bracket.
	if S[lenS-1] == 123 || S[lenS-1] == 40 || S[lenS-1] == 91 {
		return 0
	}

	stack := make([]int8, lenS/2)
	count := 0

	for _, s := range S {

		if s == 40 || s == 91 || s == 123 {

			if count+1 > (lenS)/2 {
				return 0
			}

			stack[count] = int8(s)
			count++

		} else if s == 41 {
			count--
			if count < 0 {
				return 0
			}
			if stack[count] != 40 {
				return 0
			}

		} else if s == 93 {
			count--
			if count < 0 {
				return 0
			}
			if stack[count] != 91 {
				return 0
			}

		} else if s == 125 {
			count--
			if count < 0 {
				return 0
			}
			if stack[count] != 123 {
				return 0
			}

		}

	}

	if count == 0 {
		return 1
	}

	return 0
}
