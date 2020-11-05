package cyclicrotation

// solution Rotate an array
// param A consisting of N integers
// and an integer K, returns the array A rotated K times
func solution(A []int, K int) []int {
	if K == 0 || len(A) == 0 || len(A) == 1 {
		return A
	}

	if len(A) < K {
		K = K % len(A)
	}

	lhs := A[len(A)-K:]

	return append(lhs, A[:len(A)-K]...)
}

func leftRotation(A []int, K int) []int {
	lenA:=len(A)
	
	if K == 0 || lenA == 0 || lenA == 1 || lenA == K {
		return A
	}

	if lenA < K {
		
		if K%lenA == 0 {
			return A
		} 

		K = K % len(A)
	}
	
	lhs := A[K:len(A)]

	return append(lhs, A[:K]...)
}