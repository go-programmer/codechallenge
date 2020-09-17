package maxcounter

func initialSolution(N int, A []int) []int {
	tempMaxCounter := 0
	counter := make([]int, N+1)
	counter[0] = 0

	for _, instruction := range A {

		// if A[K] = N + 1 then operation K is max counter.
		// Misjudged that condition.
		if instruction <= N {

			if tempMaxCounter == counter[0] {
				counter[instruction] = counter[0]
			}

			counter[instruction]++

			if counter[instruction] > tempMaxCounter {
				tempMaxCounter = counter[instruction]
			}

		} else {
			if tempMaxCounter > counter[0] {
				counter[0] = tempMaxCounter
			}

		}
	}

	// Set last maximum counter value to all other
	// counters whose value is less than it, counter[0].
	for key, count := range counter {
		if counter[0] > count {
			counter[key] = counter[0]
		}
	}

	return counter[1:]
}
