package solution

import "fmt"

// Returns an integer array with index as numbers
// and values representing the smallest prime factor.
func smallestPrimeFactors(n int) []int {
	F := make([]int, n+1)
	i := 2

	for i*i <= n {

		if F[i] == 0 {
			k := i * i

			for k <= n {

				if F[k] == 0 {
					F[k] = i
				}

				k += i
			}
		}
		i++
	}

	// fmt.Println(F)
	return F
}

// Return true if there are extactly two factorials of n
func twoPrimeFactors(n int, smallestPrimeFactors []int) bool {
	primesCount := 1

	for smallestPrimeFactors[n] > 0 {
		if primesCount > 2 {
			break
		}
		n /= smallestPrimeFactors[n]
		primesCount++
	}

	if primesCount == 2 {
		return true
	}

	return false
}

// Solution returns the number of semiprime numbers between
// P[i] and Q[i].
func Solution(N int, P []int, Q []int) []int {
	smallestPrimeFactors := smallestPrimeFactors(N)
	checkSemiPrime := make([]int, N+1)

	for i := 4; i < N+1; i++ {
		if twoPrimeFactors(i, smallestPrimeFactors) {

			// if _, exists := checkSemiPrime[i]; !exists {
			checkSemiPrime[i] = 1
			// }
		}
	}

	// Take prefix sum
	prefixSumSemiPrime := prefixSum(checkSemiPrime)

	fmt.Println(prefixSumSemiPrime)

	semiPrimes := make([]int, len(P))

	for i := 0; i < len(P); i++ {
		min, max := P[i], Q[i]

		if max == min {

			if prefixSumSemiPrime[max] > prefixSumSemiPrime[max-1] {
				semiPrimes[i] = 1
			} else {
				semiPrimes[i] = 0

			}

		} else if (max-min)%2 == 0 {

			if prefixSumSemiPrime[max] == prefixSumSemiPrime[min] {

				if prefixSumSemiPrime[min] > prefixSumSemiPrime[min-1] {
					semiPrimes[i] = 1
				} else {
					semiPrimes[i] = 0

				}

			} else {

				if min < 4 {
					semiPrimes[i] = prefixSumSemiPrime[max]

				} else {
					if prefixSumSemiPrime[min] > prefixSumSemiPrime[min-1] {
						semiPrimes[i] = (prefixSumSemiPrime[max] - prefixSumSemiPrime[min]) + 1

					} else {
						semiPrimes[i] = (prefixSumSemiPrime[max] - prefixSumSemiPrime[min])
					}
				}
			}

		} else {
			semiPrimes[i] = prefixSumSemiPrime[max] - prefixSumSemiPrime[min-1]
		}

	}

	return semiPrimes
}

// Return total cars passed at each point.
// Prefix sum is used for total calculation.
func prefixSum(slice []int) []int {
	lenS := len(slice)
	prefixSum := make([]int, lenS)
	prefixSum[0] = slice[0]

	for i := 1; i < lenS; i++ {
		prefixSum[i] = prefixSum[i-1] + slice[i]
	}

	return prefixSum
}
