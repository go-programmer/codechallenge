package prefixsum

func solution(slice []int) []int {
	lenS := len(slice)
	prefix := make([]int, lenS)
	prefix[0] = slice[0]

	for i := 1; i < lenS; i++ {
		prefix[i] = prefix[i-1] + slice[i]
	}

	return prefix
}
