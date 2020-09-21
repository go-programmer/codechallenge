package triangle

import (
	"sort"
)

func solution(A []int) int {

	n := len(A)

	if n < 3 {
		return 0
	}

	sort.Ints(A)

	for i := 2; i < n; i++ {
		if A[i] < A[i-1]+A[i-2] {
			return 1
		}

	}

	return 0
}
