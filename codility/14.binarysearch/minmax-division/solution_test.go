package solution

import (
	"fmt"
	"testing"
)

// func TestMakeBlocks(t *testing.T) {
// 	for k, test := range blocks[:] {
// 		fmt.Printf("k %v | test %v\n", k, test.description)

// 		if got := makeBlocks(test.sum, test.A); got != test.want {
// 			t.Fatalf("FAIL: %s \n A: %#v got %v : want %v\n",
// 				test.description, test.A, got, test.want)
// 		}

// 		// t.Logf("SUCCESS: %s", test.description)
// 	}
// }

func TestSolution(t *testing.T) {
	for k, test := range solution[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := Solution(test.K, test.M, test.A); got != test.want {
			t.Fatalf("FAIL: %s \n A: %#v got %v : want %v\n",
				test.description, test.A, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
