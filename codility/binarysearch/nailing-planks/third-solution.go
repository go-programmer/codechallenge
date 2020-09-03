package nailingplanks

/**
how many points fall in the poits(A,B)
*/

func thirdSolution(A []int, B []int, N []int) int {

	if len(N) == 1 && N[0] < A[0] {
		return -1
	}

	if len(A) == 1 {
		aValue := A[0]
		bValue := B[0]

		nails := 0
		for _, nail := range N {
			if nail >= aValue && nail <= bValue {
				nails++
			}
		}

		return nails
	}

	// check all points intersect
	if thirdDoesIntersect(A, B) == false {
		return -1
	}

	// set low and high
	low := A[0]
	high := B[len(B)-1]

	// get total nailed number
	return thirdGetNailed(low, high, N)
}

func thirdDoesIntersect(A []int, B []int) bool {
	for i := 1; i < len(A); i++ {

		if A[i] > B[i-1] {
			return false
		}
	}

	return true
}

func thirdGetNailed(low int, high int, nails []int) int {

	nailed := 0
	for _, v := range nails {

		if v >= low && v <= high {
			nailed++
		}

		if v >= high {
			break
		}

	}

	if nailed == 0 {
		return -1
	}

	return nailed
}
