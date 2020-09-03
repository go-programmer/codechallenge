package nailingplanks

/**
Expected worst-case time complexity is O((N+M)*log(M));
Expected worst-case space complexity is O(M) (not counting the storage required for input arguments).
*/

func thirdpartySolution(a []int, b []int, c []int) int {
	n := len(a)
	m := len(c)
	minNails := 1
	maxNails := m
	var mid int
	var missing bool

	total := -1

	maxCoord := m*2 + 1
	nailed := make([]int, maxCoord)

	for minNails <= maxNails {
		missing = false
		mid = (maxNails + minNails) / 2

		for i := 0; i < maxCoord; i++ {
			nailed[i] = 0
		}

		for i := 0; i < mid; i++ {
			nailed[c[i]]++
		}

		for i := 1; i < maxCoord; i++ {
			nailed[i] = nailed[i] + nailed[i-1]
		}

		for i := 0; i < n; i++ {
			if nailed[a[i]-1] == nailed[b[i]] {
				missing = true
			}
		}

		if missing {
			minNails = mid + 1
		} else {
			maxNails = mid - 1
			total = mid
		}
	}

	return total
}
