package solution

import (
	"fmt"
	"math"
)

// Returns the smallest perimeter of a
// rectangle that can be formed with area
// n
func solution(n int) int {

	if n == 1 {
		return 4
	}

	if sqr, sqrt := perfectSquare(n); sqr {
		return sqrt * 4
	}

	hd := highestDivisor(n)

	divident := n / hd

	hdSqr, _ := perfectSquare(hd)

	dSqr, _ := perfectSquare(divident)

	if hdSqr && dSqr {
		hd--
		divident--
	}

	return 2 * (hd + divident)
}

// Returns highest divisor of n
// O(√n)
func highestDivisor(n int) int {
	i := 1
	highestDivisor := 0
	divisors := make([]int, n/2)
	divisorsLen := 0

	for i*i < n {

		if n%i == 0 {
			highestDivisor = i
			divisors[divisorsLen] = i
			divisorsLen++
		}

		i++
	}
	divisors = divisors[:divisorsLen]
	fmt.Println(divisors)
	return highestDivisor
}

// Return true if n is a perfect square
func perfectSquare(n int) (bool, int) {
	sr := math.Sqrt(float64(n))

	if square := sr - math.Floor(sr); square == 0 {
		return true, int(sr)
	}

	return false, 0
}
