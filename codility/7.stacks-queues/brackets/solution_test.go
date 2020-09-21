package brackets

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	for k, test := range testCases {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.S); got != test.want {
			t.Fatalf("FAIL: %s \n S: %#v got %d : want %d\n",
				test.description, test.S, got, test.want)
		}

		t.Logf("SUCCESS: %s", test.description)
	}
}
