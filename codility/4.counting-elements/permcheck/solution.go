package permcheck

import "sort"

func solution(A []int) int {

	lenA := len(A)

	if lenA == 1 {
		if A[0] == 1 {
			return 1
		}
		return 0
	}

	sort.Ints(A)

	if A[0] != 1 || lenA != A[lenA-1] {
		return 0
	}

	count := 2
	// Skip 1st and last
	for _, val := range A[1 : lenA-1] {

		if val != count {
			return 0
		}

		count++
	}

	return 1
}
