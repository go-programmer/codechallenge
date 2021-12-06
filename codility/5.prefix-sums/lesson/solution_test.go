package lesson

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	for k, test := range testCases[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.A, test.k, test.m); got != test.want {
			t.Fatalf("FAIL: %s \n A: %#v got %d : want %d\n",
				test.description, test.A, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
