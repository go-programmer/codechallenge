package solution

import (
	"math"
)

func solution(A []int) int {

	if len(A) == 1 {
		return A[0]
	}

	// float is taken because Max()
	maxEnding, maxSlice := 0.0, float64(A[0])

	for _, a := range A {
		maxEnding = math.Max(float64(a), maxEnding+float64(a))
		maxSlice = math.Max(maxSlice, maxEnding)
	}

	return int(maxSlice)
}
