package countfactors

// Returns the number of divisors of n
// O(√n)
func solution(n int) int {
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
