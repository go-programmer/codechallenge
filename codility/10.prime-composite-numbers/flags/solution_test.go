package solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPeaks(t *testing.T) {

	for k, test := range peaks[:] {

		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := countPeaks(test.A); !reflect.DeepEqual(got, test.want) {

			t.Fatalf("FAIL: %s \n A: %v got %v : want %v\n",
				test.description, test.A, got, test.want)
		}
	}
}

func TestFlags(t *testing.T) {

	for k, test := range flags[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := countFlags(test.peaks); got != test.want {

			t.Fatalf("FAIL: %s \n n: %v got %v : want %v\n",
				test.description, test.peaks, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

func TestSolution(t *testing.T) {

	for k, test := range peakandflag[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.A); got != test.want {

			t.Fatalf("FAIL: %s \n n: %v got %v : want %v\n",
				test.description, test.A, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
