package cyclicrotation

// solutionWithValueShift rotates values in an
// array nums by k clockwise and returns the
// rotated array.
// This was the initial solution.
func solutionWithValueShift(nums []int, k int) []int {
	numsSize := len(nums)

	if k == 0 || k == numsSize || k > 100 || numsSize == 1 || numsSize == 0 || numsSize > 100 {
		return nums
	}

	if k > numsSize {
		mod := k % numsSize
		if mod == 0 {
			k = k / numsSize
		} else {
			k = mod
		}
	}

	shiftValues(nums, 0, numsSize-1)
	shiftValues(nums, 0, k-1)
	shiftValues(nums, k, numsSize-1)

	return nums
}

// shiftValues exchange values in a slice
func shiftValues(a []int, b, e int) {
	for b < e {
		t := a[b]
		a[b] = a[e]
		a[e] = t
		b++
		e--
	}
}
