package tapeequlibrium

import "math"

func solution(A []int) int {
	arraySum := 0
	currentMin := 1<<32 - 1

	for _, value := range A {
		arraySum += value
	}
	lhs := A[0]
	rhs := arraySum - lhs

	for i := 1; i < len(A); i++ {
		diff := int(math.Abs(float64(lhs) - float64(rhs)))

		if diff < currentMin {
			currentMin = diff
		}
		lhs += A[i]
		rhs -= A[i]
	}

	return currentMin
}
