package passingcars

func solution(A []int) int {
	everyCarsPassed := eachCarsPassed(A)
	lenA := len(A) - 1
	totalCarsPassed := 0

	for point, car := range A {

		// Count cars that passed from here
		if car == 0 {
			totalCarsPassed += carsPassed(everyCarsPassed, point, lenA)
		}

		// Requirement
		if totalCarsPassed > 1000000000 {
			return -1
		}
	}

	return totalCarsPassed
}

// Returns total cars that passed from point x
// till point y in the prefix sum P
func carsPassed(P []int, x int, y int) int {
	return P[y] - P[x]
}

// Return total cars passed at each point.
// Prefix sum is used for total calculation.
func eachCarsPassed(slice []int) []int {
	lenS := len(slice)
	prefixSum := make([]int, lenS)
	prefixSum[0] = slice[0]

	for i := 1; i < lenS; i++ {
		prefixSum[i] = prefixSum[i-1] + slice[i]
	}

	return prefixSum
}
