package uniquenumbers

import (
	"fmt"
	"testing"
)

func TestNailingPlanks(t *testing.T) {
	for k, test := range testCases {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := uniqueNumbers(test.A); got != test.want {
			t.Fatalf("FAIL: %s \n A: %#v got %d : want %d\n",
				test.description, test.A, got, test.want)
		}

		t.Logf("SUCCESS: %s", test.description)
	}
}
