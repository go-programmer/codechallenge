package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	// last := len(perimeter) - 1

	for k, test := range perimeter[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.n); got != test.want {

			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

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

func TestIsPerfectSquare(t *testing.T) {

	for k, test := range square[len(perimeter)-1:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got, _ := perfectSquare(test.n); got != test.want {

			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
