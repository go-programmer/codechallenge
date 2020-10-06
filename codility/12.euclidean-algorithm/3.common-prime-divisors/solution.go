package solution

// Solution returns the count of same prime factors in
// the elements of A[i] and B[i]
func Solution(A []int, B []int) int {
	count := 0

	for i := 0; i < len(A); i++ {

		if samePrimeDivisors(A[i], B[i]) {
			count++
		}
	}

	return count

}

// Returns true if N and M has same number of prime divisors.
// q >= p
func samePrimeDivisors(p, q int) bool {

	if p == q {
		return true

	}

	denom := 0
	if q%p == 0 {
		denom = p
	} else {
		denom = gcd(p, q, 1)
	}

	if denom != p {

		for p != 1 {
			pgcd := gcd(p, denom, 1)

			if pgcd == 1 {
				return false
			}
			p /= pgcd
		}
	}

	for q != 1 {
		pgcd := gcd(q, denom, 1)

		if pgcd == 1 {
			return false
		}
		q /= pgcd
	}

	return true
}

// Returns greatest common divisor of a and b
func gcd(a, b, res int) int {

	if a == b {
		return res * a

	} else if (a%2 == 0) && (b%2 == 0) {
		return gcd(a/2, b/2, 2*res)

	} else if a%2 == 0 {
		return gcd(a/2, b, res)

	} else if b%2 == 0 {
		return gcd(a, b/2, res)

	} else if a > b {
		return gcd(a-b, b, res)

	} else {
		return gcd(a, b-a, res)
	}

}
