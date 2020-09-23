package lesson

import (
	"math"
)

func solution(A []int) int {

	maxEnding, maxSlice := 0.0, 0.0

	for _, a := range A {
		maxEnding = math.Max(0, maxEnding+float64(a))
		maxSlice = math.Max(maxSlice, maxEnding)
	}

	return int(maxSlice)
}
