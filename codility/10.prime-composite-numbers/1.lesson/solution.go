package lesson

// Return true if n is prime
// O(√n)
func isPrime(n int) bool {
	i := 2

	for i*i <= n {

		if n%i == 0 {
			return false
		}
		i++
	}

	return true
}

// Returns the number of divisors of n
// O(√n)
func divisors(n int) int {
	i := 1
	result := 0

	for i*i < n {

		if n%i == 0 {
			result += 2
		}
		i++
	}

	if i*i == n {
		result++
	}

	return result
}

/*
Problem: Consider n coins aligned in a row. Each coin is showing heads at the beginning.

1,2,3,4,5,6,7,8,9

Then, n people turn over corresponding coins as follows.
Person i reverses coins with numbers that are multiples of i.

That is, person i flips coins i, 2 · i, 3 · i, . . . until
no more appropriate coins remain.

The goal is to count the number of coins showing tails.

In the above example, the final configuration is:
Tail, 2, 3, Tail, 5, 6, 7, 8, Tail

Solution O(n log n): We can simulate the results of each person reversing coins.

*/

func countTails(n int) int {
	result := 0
	coin := make([]int, n+1)
	i := 1
	for i < n+1 {
		k := i

		for k <= n {
			coin[k] = (coin[k] + 1) % 2
			k += i
		}

		result += coin[i]
		i++
	}

	return result
}

/*
There is a O(log n) solution that needs to worked out.

Solution O(log n): Notice that each coin will be turned over exactly as many times as the
number of its divisors. The coins that are reversed an odd number of times show tails, meaning
that it is sufficient to find the coins with an odd number of divisors.
We know that almost every number has a symmetric divisor (apart from divisors of the
√
form n). Thus, every number of the form k 2 has an odd number of divisors. There are
√
√
exactly � n� such numbers between 1 and n. Finding the value of � n
*/
