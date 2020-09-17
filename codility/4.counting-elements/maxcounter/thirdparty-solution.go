package maxcounter

/*
Complexity:
expected worst-case time complexity is O(N+M);
expected worst-case space complexity is O(N), beyond input storage
(not counting the storage required for input arguments).
*/

func thirdpartySolution(N int, A []int) []int {
	maxNum := N + 1
	base := 0
	currentMax := 0
	result := make([]int, N)

	for _, value := range A {
		
		if value == maxNum {
			base = currentMax
		} else {
			if result[value-1] < base {
				result[value-1] = base
			}

			result[value-1] += 1

			if currentMax < result[value-1] {
				currentMax = result[value-1]
			}
		}
	}

	for idx, value := range result {
		if value < base {
			result[idx] = base
		}
	}

	return result
}
