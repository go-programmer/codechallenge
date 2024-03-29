package maxcounter

func solution(N int, A []int) []int {

	// Incresed to make easy for comparision
	nMax := N + 1

	tempMaxCounter := 0
	counter := make([]int, N+1)

	// This is used as the last max value operation
	counter[0] = 0

	for _, instruction := range A {

		if instruction == nMax {
			counter[0] = tempMaxCounter
		} else {

			if counter[instruction] < counter[0] {
				counter[instruction] = counter[0]
			}

			counter[instruction]++

			if tempMaxCounter < counter[instruction] {
				tempMaxCounter = counter[instruction]
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
