package minavgtwoslice

import "math"

func solution(A []int) int {
	prefixSum := prefixSum(A)
	minAvg := math.MaxFloat64
	minAvgIndex := 1
	delPrevSum := 0

	for i := 0; i < len(prefixSum)-1; i++ {
		// Make a new prefix sum slice without previous
		// sum.
		newPrefixSumSlice := prefixSum[i+1:]

		sliceMinAvg := getMinAverage(newPrefixSumSlice, delPrevSum)

		if sliceMinAvg < minAvg {
			minAvg = sliceMinAvg
			minAvgIndex = i + 1
		}

		delPrevSum = prefixSum[i]
	}

	return minAvgIndex
}

// returns minimum average
func getMinAverage(prefixSum []int, delPrevSum int) float64 {
	divisor := 2
	average := math.MaxFloat64

	for _, v := range prefixSum {
		currentAvg := float64((v - delPrevSum) / divisor)

		if average > currentAvg {
			average = currentAvg
		}

		divisor++
	}

	return average
}

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
