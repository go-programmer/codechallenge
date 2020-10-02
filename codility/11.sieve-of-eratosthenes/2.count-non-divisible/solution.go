package solution

type Element struct {
	index            []int
	count            int
	nonDivisorsCount int
}

func solution(A []int) []int {
	lenA := len(A)

	// remove duplicates
	mapA := uniqueMap(A)

	for k, v := range mapA {

		// Loop through all the elements.
		// Calculate non divisors.
		for i := 0; i < lenA; i++ {

			if A[i] < k && k%A[i] != 0 {
				v.nonDivisorsCount++

			} else if A[i] > k {
				v.nonDivisorsCount++
			}
		}
	}

	for i := 0; i < lenA; i++ {
		if v, e := mapA[A[i]]; e {
			A[i] = v.nonDivisorsCount
		}
	}

	return A
}

// Returns map of elements with element property
func uniqueMap(A []int) map[int]*Element {
	uniqueNum := map[int]*Element{}

	for k, a := range A {

		if _, ok := uniqueNum[a]; !ok {
			uniqueNum[a] = &Element{[]int{k}, 1, 0}
		} else {
			uniqueNum[a].count = uniqueNum[a].count + 1
			uniqueNum[a].index = append(uniqueNum[a].index, k)
			// uniqueNum[a] = element{count, 0}
		}

	}

	return uniqueNum
}

/**
Previous logic before creating uniqueMap()

i := 1
	for i < lenA {

		nonDivisorsCount := 0

		if i != lenA-1 {

			if copyA[i] != copyA[i+1] {

				//check if the number already exists in map
				if _, exists := mapA[copyA[i]]; !exists {

					mapA[copyA[i]] = 0

					// check previous numbers divide i
					for j := i - 1; j >= 0; j-- {

						if copyA[i]%copyA[j] != 0 {
							nonDivisorsCount++
						}
					}

					mapA[copyA[i]] += lenA - (1 + i + nonDivisorsCount)
				}
			}
		}
		i++
	}

	for k, a := range A {
		A[k] = mapA[a]
	}

	return A
*/
