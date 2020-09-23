package solution

import (
	"math"
)

func solution(A []int) int {

	maxEnding, maxSlice := 0.0, float64(math.MinInt64)

	for _, a := range A {
		maxEnding = math.Max(float64(a), maxEnding+float64(a))
		maxSlice = math.Max(maxSlice, maxEnding)
	}

	return int(maxSlice)
}
