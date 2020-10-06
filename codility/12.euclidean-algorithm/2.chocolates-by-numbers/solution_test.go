package solution

import (
	"fmt"
	"testing"
)

func TestSeieve(t *testing.T) {
	for k, test := range choclates[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := Solution(test.N, test.M); got != test.want {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.N, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
