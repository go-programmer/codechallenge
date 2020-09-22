package dominator

func solution(A []int) int {
	size := 0
	value := 0

	for _, v := range A {

		if size == 0 {
			size++
			value = v

		} else {

			if value != v {
				size--
			} else {
				size++
			}
		}
	}

	if size < 1 {
		return -1
	}

	leader := -1
	count := 0

	for k, v := range A {

		if v == value {
			count++
			leader = k
		}
	}

	if count > len(A)/2 {
		return leader
	}

	return -1
}
