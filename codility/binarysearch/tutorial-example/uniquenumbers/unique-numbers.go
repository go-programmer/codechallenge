package uniquenumbers

import "sort"

// def distinct(A):
// n = len(A)
// A.sort()
// result = 1
// for k in xrange(1, n):
// if A[k] != A[k - 1]:
// result += 1
// return result

func uniqueNumbers(A []int) int {
	lenA := len(A)
	copyA := make([]int, lenA)
	copy(copyA, A)
	sort.Ints(copyA)
	result := 1

	for k := 1; k < lenA; k++ {

		if copyA[k] != copyA[k-1] {
			result += 1
		}
	}

	return result
}
