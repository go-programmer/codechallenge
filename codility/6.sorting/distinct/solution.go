package distinct

func solution(A []int) int {

	var uniqueNum = make(map[int]bool)

	for _, a := range A {

		if _, ok := uniqueNum[a]; !ok {
			uniqueNum[a] = true
		}
	}

	return len(uniqueNum)
}
