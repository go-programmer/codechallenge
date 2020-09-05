package missinginteger

import (
	"sort"
)

// solution Return smallest positive integer in the slice
func solution(A []int) int {
	sort.Ints(A)
	lenA := len(A)
	lastIndex := lenA - 1

	if A[lastIndex] <= 0 || A[0] > 1 {
		return 1
	}

	if lenA == 1 && A[0] == 1 {
		return 2
	}

	i := sort.Search(lenA, func(i int) bool { return A[i] >= 1 })
	if i < lenA && A[i] != 1 {
		return 1
	}

	lastValue := 1

	for i < lastIndex {

		if A[i] == A[i+1] {
			lastValue = A[i] + 1
		} else {

			nextVal := A[i] + 1

			if A[i] != nextVal && nextVal != A[i+1] {
				lastValue = nextVal
				break
			}

		}

		i++
	}

	// Above for loop will not check for last value
	if i == lastIndex {

		// If second last and last value are equal
		if A[i-1] == A[i] {
			lastValue = A[i] + 1
		}

		// If last value is one more than last value
		if A[i-1]+1 == A[i] {
			lastValue = A[i] + 1
		}

	}

	if lastValue >= 1000000 {
		return 1000000
	}

	return lastValue

}

// Pseudocode:
// sort
// if the max is -ve or 0, then return 1
// remove duplicate - No need.
// binary search for index whose value is 1
// then compare next smallest number starting from that index
// 		if next index is not equal to or greater than A[index} + 1,
// 		return A[index] + 1
