package nailingplanks

/**
how many points fall in the poits(A,B)
*/

func secondSolution(A []int, B []int, N []int) int {

	// check all points intersect
	if doesIntersectSecond(A, B) == false {
		return -1
	}

	// set low and high
	low := A[0]
	high := B[len(B)-1]

	// get total nailed number
	return getNailedSecond(low, high, N)
}

func doesIntersectSecond(A []int, B []int) bool {

	for i := 1; i < len(A); i++ {

		if A[i] > B[i-1] {
			return false
		}
	}

	return true
}

func getNailedSecond(low int, high int, nails []int) int {

	if len(nails) == 1 && nails[0] < low {
		return -1
	}

	nailed := 0
	for _, v := range nails {

		if v >= low && v <= high {
			nailed++
		}

		if v >= high {
			break
		}

	}

	return nailed
}

// func solution(A []int, B []int, N []int) int {

// 	// np := 0 // nailed planks count
// 	lN := len(N)
// 	lB := len(B)
// 	low := A[0]
// 	high := B[0]

// 	i, j := 0, 0

// 	for i < lN-1 {
// 		for j < high {

// 			// If N is between low and high
// 			if N[j] >= low && N[j] <= high {
// 				j++
// 			}

// 			if low == high {
// 				i++
// 				break
// 			}
// 		}
// 		// Two points do not intersect
// 		if A[i+1] > B[i] {
// 			return -1
// 		}

// 		high = B[i]
// 	}

// 	return -1
// 	/*
// 		Put all the points in the same line,
// 		we need to check if the nails can be fitted in.
// 	*/

// }

func solutions(A []int, B []int, N []int) int {
	low := A[0]
	high := B[0]
	np := 0 // nailed planks count
	lN := len(N)
	lA := len(A)
	i, j := 0, 1

	for i < lN-1 {

		if N[i] >= low && N[i] <= high {
			np++
			low = N[i]

			if low == high {
				// If there is the two coordinates does
				// not have an intersection point
				if A[i] >= A[i+1] && A[i] <= B[i+1] {
					return -1
				}
				j++
				high = B[j]

			}

		} else {
			j++
			high = B[j]
			i++
		}

		if j == lA-1 {
			return np
		}
	}

	return -1

}

/**
# Thought process 2

If nail is between low nail and high nail,
increase nailed planks!

low = A[0]
high = B[0]
np = 0 // nailed planks count
lN = len(N)
lA = len(A)
highIndex = 0

loop i = 0;j=0; i < lN - 1 {
//within this loop all planks must be nailed.
	// (1,4)
	// If N is between low and high, inclusive!
	if N[i] >= low && N[i] <= high {
		np++
		low= N[i]

		if low == high {
			if A[i] >= A[i+1] && A[i] <= B[i+1] {
					return -1
			}
			high = B[j++]

		}

	}

	if j == len(B)-1 {
		return np
	}

	return -1

}

// P{(1,4), (4,5), (5,9), (8,10)} ; N{1,2,3,4,6,7,10}
l=1;h=4
//Till  1,2,3,4; np=4;
l==h {
	B[i] >= B[i+1]
}

*/
