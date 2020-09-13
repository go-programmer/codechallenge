package invertedtriangle

import (
	"fmt"
	"testing"
)

func TestPrintSolution(t *testing.T) {
	want := `*********** *********   *******     *****       ***         *     `

	if got := solution(11); got != want {
		t.Fatalf("FAIL")
	}

	// t.Logf("SUCCESS: %s", test.description)

}

func TestSolution(t *testing.T) {
	for k, test := range testCases {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.n); got != test.want {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		t.Logf("SUCCESS: %s", test.description)
	}
}
