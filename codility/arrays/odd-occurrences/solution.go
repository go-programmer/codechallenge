package oddoccurences

import "sort"

// solution finds and returns the
// odd number in an array.
// Solution analysis result at solution.png
func solution(A []int) (oddOne int) {

	sort.Ints(A)

	lenA := len(A)
	i := 0
	for i < lenA-1 {

		if A[i] != A[i+1] {
			return A[i]
		}

		i = i + 2
	}

	// If lenA is odd and value is at
	// the last index.
	return A[lenA-1]
}
