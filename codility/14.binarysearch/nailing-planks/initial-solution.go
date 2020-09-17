package nailingplanks

func initialSolution(A []int, B []int, N []int) int {
	/*
		Planks P[A, B] : [1, 4], [4, 5], [5, 9] and [8, 10]
		Nails N : 9, 4, 6, 7, 10, 2

		First J nails are used.
		Nail must be A[K] ≤ C[I] ≤ B[K]
		Return number of nails.
		If it is not possible to nail all the planks, the function should return −1.

		Nail must be between Current Last point and Next Last point.
		Or simply nail must be in the range of both planks.

		Check if the the first plank can be nailed?
		N must be between:
			P1[A] and P2[B]
			P1[B] and P2[A]
				TRUE: Increase N count

		Check for next, from with next nail



	*/
	nailed := 0
	nails := len(N)
	plankA := len(A)

	// nails are less than planks
	if nails-1 < plankA {
		return -1
	}

	plank, nail := 0, 0

	for plank < plankA-1 {

		// check if a plank can be nailed.
		if N[nail] >= A[plank] && N[nail] <= B[plank+1] {

			if B[plank] >= A[plank+1] && B[plank] <= B[plank+1] {
				nailed++
				plank++
			}
		}
		// check if there are more nails left
		if nail+1 > nails-1 {
			return -1
		}

		nail++
	}

	//check all planks are nailed
	if nailed == plank {
		return nailed + 1
	}

	return -1
}
