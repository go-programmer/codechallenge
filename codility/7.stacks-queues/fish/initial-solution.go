package fish

func initialSolution(A []int, B []int) int {

	count := 0
	upIndex := 0
	eat := 0

	for i := 0; i < len(A); i++ {
		// fish comes to the river
		count++

		if B[i] == 1 {
			upIndex = i
			eat = 1

		} else {
			// check if gets eaten or not
			if eat == 1 {

				// if bigger fish was swimming down,
				// then the fish gets eaten. Else this
				// fish kills the eating type fish.
				if A[upIndex] > A[i] {

				} else {
					eat = 0
				}

				count--
			}

		}

	}

	return count
}
