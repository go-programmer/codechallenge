package suffixsum

func solution(slice []int) []int {
	lenS := len(slice)
	suffixsum := make([]int, lenS)
	suffixsum[lenS-1] = slice[lenS-1]

	for i := lenS - 1; i > -1; i-- {
		suffixsum[i] = suffixsum[i-1] + slice[i]
	}

	return suffixsum
}

func solutionSameSlice(array []int) []int {

	for i := len(array) - 2; i > -1; i-- {
		array[i] = array[i+1] + array[i]
	}

	return array
}
