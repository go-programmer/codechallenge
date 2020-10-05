package lesson

import "fmt"

// Returns array of n numbers from 0.
// Index represents the number and
// the value 0 means the index number is prime,
// 1 means composite.

// The logic is to represent 0 for prime in contrast
// to the tutorial.

// This function implments sieve concept.
func markPrime(n int) []int {
	sieve := make([]int, n+1)

	sieve[0] = 1
	sieve[1] = 1
	i := 2

	for i*i <= n {

		if sieve[i] == 0 {
			k := i * i

			for k <= n {
				sieve[k] = 1
				k += i
			}
		}

		i++
	}

	fmt.Println(sieve)

	return sieve
}

// Returns an integer array with index as numbers
// and values representing the smallest prime factor.
// This is the exact implementation as above sieve().
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

	fmt.Println(F)
	return F
}

// Return the factorial of n.
// Complexity: O(log n)
func primeFactors(n int) []int {
	primeFactors := []int{}

	n = 9

	F := smallestPrimeFactors(n)

	for F[n] > 0 {
		primeFactors = append(primeFactors, F[n])

		n /= F[n]
	}

	primeFactors = append(primeFactors, n)

	return primeFactors
}
