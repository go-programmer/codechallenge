package solution

import (
	"sort"
)

// Returns the peaks array and the index of last peak
// of the mountains.
func countPeaks(A []int, lenA int) []int {
	peaks := make([]int, lenA/2)
	lastIndex := 0

	for i := 1; i < lenA-1; i++ {

		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks[lastIndex] = i
			lastIndex++
		}

	}

	return peaks[:lastIndex]
}

// Returns highest divisor of n
// O(√n)
func factorials(n int) ([]int, int) {
	i := 1
	divisors := make([]int, n)
	divisorsLen := 0

	for i*i <= n {

		if n%i == 0 {
			divisors[divisorsLen] = i
			divisorsLen++
			divisors[divisorsLen] = n / i
			divisorsLen++
		}

		i++
	}

	// if i*i == n {
	// 	result++
	// }

	return divisors[:divisorsLen], divisorsLen
}

func solution(A []int) int {
	lenA := len(A)

	if lenA < 3 {
		return 0
	}

	peaks := countPeaks(A, lenA)
	peaksLen := len(peaks)

	// check if there are no peak return 0
	if peaksLen == 0 {
		return 0
	} else if peaksLen == 1 {
		return 1
	}

	// Check for prime number
	factorials, factorialsLen := factorials(lenA)
	sort.Ints(factorials)

	// if lenA is prime
	if factorialsLen == 2 && peaksLen > 1 {
		return 1
	}

	// Divide A into blocks containing equal number of elements.
	// Each block must contain at least on peak.
	// The number of peaks determine if the maximum possible blocks.
	// Block of size 1 and two are not possible.

	maxBlocksIndex := 1
	// Block of size two is not possible.
	// So skip is second index value is 2.
	if factorials[1] == 2 {
		maxBlocksIndex = 2
	}

	blocks := factorials[factorialsLen-maxBlocksIndex]

	// Skip the last one
	for blocks > 1 {

		if peaksLen >= blocks {

			if makeBlocks(lenA/blocks, peaks, peaksLen, lenA) {
				return blocks
			}
		}

		maxBlocksIndex++
		blocks = factorials[factorialsLen-maxBlocksIndex]
	}

	return 1
}

// Returns true if blocks of size b is possible.
func makeBlocks(b int, peaks []int, peaksLen int, lenA int) bool {

	peakIndex := 0
	// As we have stored the indexes of peaks, we can use it check if a peak exists in the block.
	j := 0
	for j < lenA {
		hasPeak := false
		min, max := j, j+b

		// The block must contain a peak
		for k := min; k < max; k++ {

			if peaks[peakIndex] >= min && peaks[peakIndex] < max {
				peakIndex++
				hasPeak = true
			}

			// if all peaks are finished but more blocks
			// are left.
			if peakIndex == peaksLen && j < lenA {

				if hasPeak == true {
					return true
				}

				return false
			}
		}

		// if no peak exists in the block
		if hasPeak == false {
			return false
		}

		// Check next block
		j += b
	}

	return false
}

/*
Previous logic

for i := 1; i < factorialsLen; i++ {
		factorial := factorials[i]
		secondDivisor := lenA / factorial
		blocks := make([]int, 2)

		if peaksLen >= secondDivisor {
			blocks[0] = secondDivisor
			blocks[1] = factorial

		} else if peaksLen >= factorial {
			blocks[0] = factorial
			blocks = blocks[0:1]
		}

		// check if each block in the blocks of the size
		// blockSizes contains at least one peak
		for _, v := range blocks {
			blockSize := lenA / v

			// checkAllPeaks := peaksLen
			// for checkAllPeaks > 0 {
			peakIndex := 0
			// blocksCreated := 0

			// As we have stored the indexes of peaks, we can use it check if a peak exists in the block.
			j := 0
			for j < lenA {
				hasPeak := false
				min, max := j, j+blockSize

				// how many peaks falls under min and max
				// at least one is needed.
				for k := min; k < max; k++ {

					if peaks[peakIndex] >= min && peaks[peakIndex] < max {
						peakIndex++
						hasPeak = true

						// All blocks contains at least one peak
						// And check this is the last block
						if peakIndex >= peaksLen && j >= lenA-blockSize {

							if maxBlockSize < blockSize {
								maxBlockSize = blockSize
								break
							}

							// return blocksCreated + 1
						}
					}
				}

				// if no peak exists in the block
				if hasPeak == false {
					break
				}

				// if all peaks are finished but more blocks
				// are left.
				if peakIndex > peaksLen && j < lenA {
					break
				}

				// Check next block
				j += blockSize
				// blocksCreated++
				// }

				// checkAllPeaks--
			}

		}
	}

	return maxBlockSize
*/
