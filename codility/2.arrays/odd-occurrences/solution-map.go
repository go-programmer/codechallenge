package oddoccurences

// solutionWithHashMap finds and returns the
// odd number in an array.
// This was the initial solution.
// Solution analysis result at solution-map.png
func solutionWithHashMap(A []int) int {
	var histogram = make(map[int]int)
	countOccurances(A, histogram)

	k := 1
	value := 0
	for k, value = range histogram {

		if isOdd(value) {
			break
		}
	}

	return k
}

// countOccurances counts the occurnace of numbers
// in an array A and stores the total count in a
// map occuranceMap.
func countOccurances(A []int, occuranceMap map[int]int) {
	for _, value := range A {
		occuranceMap[value]++
	}
}

// isOdd returns true if the number n is odd,
// else false.
func isOdd(n int) bool {
	return n%2 == 1
}
