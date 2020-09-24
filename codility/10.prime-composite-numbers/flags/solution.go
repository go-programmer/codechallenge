package solution

// Return the flags set in the peaks.
func solution(A []int) int {

	if len(A) < 3 {
		return 0
	}

	peaks := countPeaks(A)

	if len(peaks) == 0 {
		return 0
	}

	return countFlags(peaks)
}

// Returns the peaks array and the index of last peak
// of the mountains.
func countPeaks(mountains []int) (peaks []int) {
	lastIndex := 0
	// Assuming the maximum number of the peaks
	// will be 1/3 of total array.
	peaks = make([]int, len(mountains))

	for i := 1; i < len(mountains)-1; i++ {

		if mountains[i] > mountains[i-1] && mountains[i] > mountains[i+1] {
			peaks[lastIndex] = i
			lastIndex++
		}
	}

	peaks = peaks[0:lastIndex]
	return
}

func countFlags(peaks []int) int {
	peakLen := len(peaks) - 1
	flag := 1

	for i := 0; i < peakLen && flag <= peakLen; i++ {

		if peaks[i+1]-peaks[i] >= peakLen {
			flag++

		} else if peaks[i+2]-peaks[i] >= peakLen {
			flag++
			i++
		}
	}

	return flag
}
