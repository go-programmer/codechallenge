package reversestring

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	for k, test := range testCases {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.word); got != test.want {
			t.Errorf("Reverse - %s: got:want %v:%v\n", test.description, got, test.want)

		}

		t.Logf("SUCCESS: %s", test.description)
	}
}

func TestLinearSolution(t *testing.T) {
	for k, test := range testCases[1:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := linearSolution(test.word); got != test.want {

			t.Fatalf("\nFAIL: %s n: %#v got %v : want %v\n",
				test.description, test.word, got, test.want)

		}

		t.Logf("SUCCESS: %s", test.description)
	}
}
