package lesson

import (
	"fmt"
	"math"
)

// Return prefix sum of slice.
func prefixSum(slice []int) []int {
	lenS := len(slice)
	prefixSum := make([]int, lenS)
	prefixSum[0] = slice[0]

	for i := 1; i < lenS; i++ {
		prefixSum[i] = prefixSum[i-1] + slice[i]
	}
	return prefixSum
}

func countTotal(P []int, x int, y int) int {
	return P[y+1] - P[x]
}

func solution(A []int, k int, m int) int {
	// n := len(A)
	result := 0
	// pref := prefixSum(A)
	// rangeLimit := m
	// leftPos := 0
	// rightPos := 0

	// if m > k {
	// 	rangeLimit = k
	// }

	// a, b := 45, -42
	fmt.Println(math.MinInt8) // Output: -42

	// 	fmt.Print(min)

	// 	for i := 0; i < rangeLimit+1; i++ {
	// 		leftPos = k - i
	// 		rightPos = math.MinInt8(n-1, math.Max(k, k+m-2*i))
	// 		result = math.Max(result, countTotal(pref, leftPos, rightPos))
	// 	}

	// 	if m+1 > n-k {
	// 		rangeLimit = n - k

	// 	} else {
	// 		rangeLimit = m + 1
	// 	}

	// 	for p := 0; p < rangeLimit; p++ {
	// 		rightPos = k + p
	// 		leftPos = math.MaxInt8(0, math.MinInt8(k, k-(m-2*p)))
	// 		result = math.MaxInt8(result, countTotal(pref, leftPos, rightPos))
	// 	}

	return result
}
