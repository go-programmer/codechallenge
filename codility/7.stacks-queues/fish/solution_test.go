package fish

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	for k, test := range testCases[4:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.A, test.B); got != test.want {
			t.Fatalf("FAIL: %s \n A: %#v, B:%v got %d : want %d\n",
				test.description, test.A, test.B, got, test.want)
		}

		t.Logf("SUCCESS: %s", test.description)
	}
}
