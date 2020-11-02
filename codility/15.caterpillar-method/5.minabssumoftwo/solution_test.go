package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	for k, test := range numbers[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.M, test.A); got != test.want {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.A, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
