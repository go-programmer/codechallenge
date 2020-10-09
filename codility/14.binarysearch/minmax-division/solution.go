package solution

import (
	"sort"
)

// Solution returns maximum number of blocks
func Solution(K int, M int, A []int) int {
	lenA := len(A)

	// Take max number
	B := make([]int, lenA)
	copy(B, A)
	sort.Ints(B)
	lowerBound := B[lenA-1]

	if K >= len(A) {
		return int(lowerBound)
	}

	upperBound := 0
	for _, value := range A {
		upperBound += value
	}

	if K == 1 {
		return int(upperBound)
	}

	for lowerBound <= upperBound {
		mid := (lowerBound + upperBound) / 2

		if makeBlocks(A, K, mid) {
			upperBound = mid - 1
		} else {
			lowerBound = mid + 1
		}
	}

	return int(lowerBound)
}

// Returns true if array can be divided into blocks of size
func makeBlocks(A []int, K int, size int) bool {
	blockSum := 0
	blockCnt := 0

	for _, element := range A {

		if blockSum+element > size {
			blockSum = element
			blockCnt++

		} else {
			blockSum += element
		}

		if blockCnt >= K {
			return false
		}
	}

	return true
}
