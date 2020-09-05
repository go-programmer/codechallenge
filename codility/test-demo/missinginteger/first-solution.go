package missinginteger

import (
	"sort"
)

func firstSolution(A []int) int {

	sort.Ints(A)
	minVal := 1

	if len(A) == 1 && (A[0] >= 1000000 || A[0] < 0) {
		return 1
	}

	if A[len(A)-1] <= 0 {
		return 1
	}

	lenA := len(A) - 2
	i := 0
	for i <= lenA {
		if (A[i] >= 0) && (A[i+1]-A[i] > 1) {
			return A[i] + 1
		}
		i++
	}

	if (i-1 == lenA) && (A[i] > minVal) {
		return A[i] + 1
	}

	return minVal
}
