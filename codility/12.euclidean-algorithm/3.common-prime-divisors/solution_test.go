package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {

	for k, test := range countPrimeFactors[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := Solution(test.A, test.B); got != test.want {

			t.Fatalf("FAIL: \n %#v got %v : want %v\n",
				test.description, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
func TestSamePrimeDivisors(t *testing.T) {

	for k, test := range primeDivisors[2:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := samePrimeDivisors(test.N, test.M); got != test.want {

			t.Fatalf("FAIL: \n %#v got %v : want %v\n",
				test.description, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

// Logic Verification

func TestHighestDivisor(t *testing.T) {

	for k, test := range divisor[len(divisor)-1:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := highestDivisor(test.n); got != test.want {

			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

// Returns highest divisor of n
// O(√n)
func highestDivisor(n int) int {
	i := 1
	divisors := make([]int, n/2)
	divisorsLen := 0

	for i*i < n {

		if n%i == 0 {
			divisors[divisorsLen] = i
			divisorsLen++
		}

		i++
	}
	divisors = divisors[:divisorsLen]
	// fmt.Println(divisors)
	return divisors[divisorsLen-1]
}

func TestThirdparyLogic(t *testing.T) {
	p := 15
	q := 75
	denom := gcd(p, q, 1)

	for p != 1 {
		pgcd := gcd(p, denom, 1)

		if pgcd == 1 {
			t.Error("Fail")
		}
		p /= pgcd
	}

	for q != 1 {
		pgcd := gcd(q, denom, 1)

		if pgcd == 1 {
			t.Error("Fail")
		}
		q /= pgcd
	}

}

func gcdbyDivision(p, q int) int {

	if q == 0 {
		return p
	}

	return gcdbyDivision(q, p%q)
}
