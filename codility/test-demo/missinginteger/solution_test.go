package missinginteger

import "sort"

func Solution(A []int) int {
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
		// fmt.Printf("%d not found in %v\n", 1, A)
	}

	// if isEven(lenA) {
	// 	lastIndex = lastIndex - 1
	// }
	lastValue := 1

	for i < lastIndex {

		if A[i] == A[i+1] {
			lastValue = A[i] + 1
		} else {
			nextVal := A[i] + 1

			if A[i] != nextVal {

				if nextVal != A[i+1] {

					lastValue = nextVal
					break
				}
			}
		}

		i++
	}

	if i == lastIndex {

		// Check if second last and last value are equal
		if A[i-1] == A[i] {
			lastValue = A[i] + 1
		}

		// if last value is one more than last value
		if A[i-1]+1 == A[i] {
			lastValue = A[i] + 1
		}

	}

	if lastValue >= 1000000 {
		return 1000000
	}

	return lastValue
}
