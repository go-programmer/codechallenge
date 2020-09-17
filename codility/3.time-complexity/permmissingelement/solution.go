package permmissingelement

import (
	"sort"
)

func solution(A []int) int {
	lenA := len(A)

	if lenA == 0 {
		return 1
	}

	lastIndex := lenA - 1
	sort.Ints(A)

	// If the slice of all negative numbers
	if A[lastIndex] <= 0 {
		return 1
	}

	if lenA == 1 && A[0] == 1 {
		return 2
	}

	// The number must start from one.
	// If 1 does not exists return 1.
	i := sort.Search(lenA, func(i int) bool { return A[i] >= 1 })
	if i < lenA && A[i] != 1 {
		return 1
	}

	// The next number must be 1 more than start,
	// or equal to start.
	start := A[i]
	for i < lastIndex {
		i++

		if start != A[i] {
			start++

			if start != A[i] {
				break
			}
		}
	}

	// If last two values are equal
	if start == A[i] {
		start++
	}

	// Requirement
	if start >= 1000000 {
		return 1000000
	}

	return start
}
