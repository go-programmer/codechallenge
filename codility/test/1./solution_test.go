package lesson

import (
	"fmt"
	
	"testing"
)

func TestSolution(t *testing.T) {
	for k, test := range sieve[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := Solution(test.S); got != test.want {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.S, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}