package solution

// Returns the number of divisors of n
// O(√n)
func divisor(n int) int {
	i := 1
	result := 0

	for i*i < n {

		if n%i == 0 {
			result += 2
		}
		i++
	}

	if i*i == n {
		result++
	}

	return result
}

func initalSolution(A []int) int {
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
	divisor := divisor(lenA)

	if divisor == 2 && peaksLen > 1 {
		return 1
	} else if divisor == 2 && peaksLen == 0 {
		return 0
	}

	// Divide A into blocks containing equal number of elements.
	// Each block must contain at least on peak.
	// The number of peaks determine if the maximum possible blocks.

	for peaksLen > 0 {

		// To divide A into equal number of blocks with at leat one
		// peak in the block, lenPeak must be a factor of lenA.
		if lenA%peaksLen == 0 {

			// lets do this in second refactoring
			// first lets check for larger block
			// lenA%block

			// b - 0-3,4-7,8-11
			// p - 3,5,10

			blockSize := lenA / peaksLen
			peakIndex := 0
			blocksCreated := 0

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
							return blocksCreated + 1
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
				blocksCreated++
			}

			// How to know if all the blocks has a peak?
			// We have looped into every block and checked if the
			// peak exists in each.

		}

		peaksLen--
	}

	return 0
}
