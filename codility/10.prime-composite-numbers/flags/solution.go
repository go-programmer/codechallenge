package solution

import (
	"math"
)

type peak struct {
	peaks              []int
	totalPeaksDistance int
	minPeakDistance    int
	maxPeakDistance    int
}

// Returns the peaks array and the index of last peak
// of the mountains.
func countPeaks(mountains []int) ([]int, map[int]int, int, int, int, int) {
	flag := 0
	lastIndex := 0
	// Assuming the maximum number of the peaks
	// will be 1/3 of total array.
	peaks := make([]int, len(mountains))
	distance := make([]int, len(mountains))
	distanceCount := make(map[int]int)
	minPeakDistance := math.MaxInt8
	maxPeakDistance := math.MinInt8
	totalPeaksDistance := 0
	j := 0

	for i := 1; i < len(mountains)-1; i++ {

		if mountains[i] > mountains[i-1] && mountains[i] > mountains[i+1] {
			peaks[lastIndex] = i

			if peaks[1] > 2 {
				distance[j] = int(math.Abs(float64(peaks[lastIndex-1]) - float64(peaks[lastIndex])))
				totalPeaksDistance += distance[j]

				if _, keyExists := distanceCount[distance[j]]; keyExists {
					distanceCount[distance[j]]++

				} else {
					distanceCount[distance[j]] = 1
				}

				if distance[j] < minPeakDistance {
					minPeakDistance = distance[j]
				}

				if distance[j] > maxPeakDistance {
					maxPeakDistance = distance[j]
				}

				j++
			}

			lastIndex++
		}
	}
	peaks = peaks[0:lastIndex]
	distance = distance[0:j]

	if len(peaks) == 1 {
		flag = 1
	}

	// calculate flag

	// return
	return peaks[0:lastIndex], distanceCount, flag, totalPeaksDistance, minPeakDistance, maxPeakDistance
}

func countFlags(peaks []int) int {

	// get total
	minPeakDistance := math.MaxInt8
	maxPeakDistance := math.MinInt8
	totalPeaksDistance := 0
	peakDistancesCount := make([]int, len(peaks)/2)

	for i := 0; i < len(peaks)-1; i++ {
		distance := peaks[i+1] - peaks[i]

		peakDistancesCount[distance]++
		totalPeaksDistance += distance

		if distance < minPeakDistance {
			minPeakDistance = distance
		}

		if distance > maxPeakDistance {
			maxPeakDistance = distance
		}

	}

	if minPeakDistance == maxPeakDistance {
		return len(peaks)
	}

	return 1
}

func solution(mountains []int) int {
	lastIndex := 0
	// Assuming the maximum number of the peaks
	// will be 1/3 of total array.
	peaks := make([]int, len(mountains))
	distance := make([]int, len(mountains))
	distanceCount := make(map[int]int)
	minPeakDistance := math.MaxInt8
	maxPeakDistance := math.MinInt8
	totalPeaksDistance := 0
	j := 0

	for i := 1; i < len(mountains)-1; i++ {

		if mountains[i] > mountains[i-1] && mountains[i] > mountains[i+1] {
			peaks[lastIndex] = i

			if peaks[1] > 2 {
				distance[j] = int(math.Abs(float64(peaks[lastIndex-1]) - float64(peaks[lastIndex])))
				totalPeaksDistance += distance[j]

				if _, keyExists := distanceCount[distance[j]]; keyExists {
					distanceCount[distance[j]]++

				} else {
					distanceCount[distance[j]] = 1
				}

				if distance[j] < minPeakDistance {
					minPeakDistance = distance[j]
				}

				if distance[j] > maxPeakDistance {
					maxPeakDistance = distance[j]
				}

				j++
			}

			lastIndex++
		}
	}
	peaks = peaks[0:lastIndex]
	distance = distance[0:j]

	if len(peaks) == 1 {
		return 1
	}

	if len(peaks) == 2 {
		return 2
	}

	for i := len(peaks); i > 0; i-- {

		if i*i <= totalPeaksDistance {
			flags := 1
			prevSum := 0

			for j := 0; j < len(distance); j++ {

				if prevSum+distance[j] >= i {
					flags++
					prevSum = 0
				} else {
					prevSum += distance[j]
				}
			}

			if flags >= i {
				return i
			}
		}
	}

	return 0
}
