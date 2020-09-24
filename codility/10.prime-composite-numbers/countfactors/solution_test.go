package countfactors

import (
	"fmt"
	"testing"
)

func TestDivisor(t *testing.T) {

	for k, test := range factors[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.n); got != test.want {

			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
