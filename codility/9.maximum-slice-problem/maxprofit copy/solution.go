package solution

import (
	"math"
)

func solution(A []int) int {
	minPrice := math.MaxFloat64
	maxProfit := 0.0

	for _, value := range A {
		minPrice = math.Min(minPrice, float64(value))
		maxProfit = math.Max(float64(float64(value)-minPrice), maxProfit)
	}

	return int(maxProfit)
}

func trivialSolution(A []int) int {
	max := 0.0

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {

			if A[j] > A[i] {
				max = math.Max(max, float64(A[j]-A[i]))
			}
		}
	}

	if max > 0 {
		return int(max)
	}

	return 0

}
