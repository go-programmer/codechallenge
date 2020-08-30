package oddoccurences

import "sort"

// oddOccurences search if odd occurance of an int
// Implementation with simple for loop
// Time Complexity - O(N) or O(N*log(N))
// View result at Solution-..png
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

	// all but one of the values in A occur an even number of times.
	return A[lenA-1]
}

// Add values to a histogram and check which one is left with an odd number
// Time Complexity - O(N) or O(N*log(N))
// View result at SolutionWithHasMap-..png
func solutionWithHashMap(A []int) int {
	var histogram = make(map[int]int)
	countIntsInA(A, histogram)

	k := 1
	value := 0
	for k, value = range histogram {

		if isOdd(value) {
			break
		}
	}

	return k
}

func countIntsInA(A []int, histogram map[int]int) {
	for _, value := range A {
		histogram[value]++
	}
}

func isOdd(n int) bool {
	return n%2 == 1
}
