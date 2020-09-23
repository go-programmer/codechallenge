package solution

import "math"

func solution(A []int) int {
	maxSliceSum := 1 - (1<<32 - 1)
	N := len(A) - 2

	if len(A) < 4 {
		return 0
	}

	leftSum := make([]int, N)
	rightSum := make([]int, N)

	for i := 0; i < N-1; i++ {
		leftValue := A[i+1]
		rightValue := A[N-i]

		leftSum[i+1] = int(math.Max(0, float64(leftValue)+float64(leftSum[i])))
		rightSum[N-i-2] = int(math.Max(0, float64(rightValue)+float64(rightSum[N-i-1])))
	}

	for i := 0; i < N; i++ {
		maxSliceSum = int(math.Max(float64(maxSliceSum), float64(rightSum[i])+float64(leftSum[i])))
	}

	return maxSliceSum
}
