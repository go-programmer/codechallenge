package tapeequlibrium

import "math"

/*

Complexity:
Expected worst-case time complexity is O(N);

Expected worst-case space complexity is O(N),
beyond input storage (not counting the storage
required for input arguments).

*/

func TapeEquilibrium(A []int) int {
	arraySum := 0
	currentMin := 1<<32 - 1

	for _, value := range A {
		arraySum += value
	}
	lhs := A[0]
	rhs := arraySum - lhs

	for i := 1; i < len(A); i++ {
		diff := int(math.Abs(float64(lhs) - float64(rhs)))

		if diff < currentMin {
			currentMin = diff
		}
		lhs += A[i]
		rhs -= A[i]
	}

	return currentMin
}
