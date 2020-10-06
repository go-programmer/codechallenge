package solution

import (
	"fmt"
	"testing"
)

func TestSeieve(t *testing.T) {
	for k, test := range gcdData[4:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := gcd(test.a, test.b, 1); got != test.want {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.a, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
