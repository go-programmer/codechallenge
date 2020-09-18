package countdiv

import "math"

func solution(A int, B int, K int) int {

	divisor := int(math.Floor(float64(B/K))) - int(math.Floor(float64((A)/K)))

	if A%K == 0 {
		return divisor + 1
	}

	return divisor
}
