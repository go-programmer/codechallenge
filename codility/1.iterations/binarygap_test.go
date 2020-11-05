package iterations

import (
	"fmt"
	"testing"
)

// 9 has binary representation 1001 and contains a binary gap of length 2

var testcase = []struct {
	description string
	N           int
	want        int
	err         string
}{
	{
		description: "Provided example",
		N:           9,
		want:        2,
		err:         "",
	},
}

func TestSolution(t *testing.T) {
	for k, test := range testcase[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := longestBinaryGap(test.N); got != test.want {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.N, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
