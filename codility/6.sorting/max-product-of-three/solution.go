package distinct

import (
	"sort"
)

func solution(A []int) int {

	lenA := len(A)

	if lenA == 3 {
		return A[0] * A[1] * A[2]
	}

	sort.Ints(A)

	if A[lenA-1] < 0 {
		return A[lenA-1] * A[lenA-2] * A[lenA-3]
	}

	if A[0]*A[1] > A[lenA-2]*A[lenA-3] {
		return A[0] * A[1] * A[lenA-1]
	}

	return A[lenA-1] * A[lenA-2] * A[lenA-3]
}
