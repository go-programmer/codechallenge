package solution

import (
	"fmt"
	"testing"
)

func TestFinbonacci(t *testing.T) {
	for k, test := range numbers[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := fibonacciDynamic(test.n); got != test.want {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
