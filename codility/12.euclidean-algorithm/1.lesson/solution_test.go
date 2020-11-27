package solution

import (
	"fmt"
	"testing"
)

func TestSeieve(t *testing.T) {
	for k, test := range gcdData[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := gcdNew(test.a, test.b); got != test.want {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.a, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
