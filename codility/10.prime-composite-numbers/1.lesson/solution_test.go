package lesson

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	for k, test := range primality[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := isPrime(test.n); got != test.want {

			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

func TestDivisor(t *testing.T) {
	for k, test := range divisor[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := divisors(test.n); got != test.want {

			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

func TestCountTails(t *testing.T) {
	for k, test := range coins[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := countTails(test.n); got != test.want {

			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
