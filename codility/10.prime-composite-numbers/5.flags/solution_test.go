package solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPeaks(t *testing.T) {

	for k, test := range peaks[:] {

		fmt.Printf("k %v | test %v\n", k, test.description)
		actualPeak, actualDistances, _, _, _, _ := countPeaks(test.A)

		if !reflect.DeepEqual(actualPeak, test.expectedPeaks) {

			t.Fatalf("FAIL: %s \n A: %v got %v : want %v\n",
				test.description, test.A, actualPeak, test.expectedPeaks)
		}

		if !reflect.DeepEqual(actualDistances, test.expectedDistances) {

			t.Fatalf("FAIL: %s \n A: %v got %v : want %v\n",
				test.description, test.A, actualDistances, test.expectedDistances)
		}

	}
}

func TestFlags(t *testing.T) {

	for k, test := range flags[2:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := countFlags(test.peaks); got != test.want {

			t.Fatalf("FAIL: %s \n n: %v got %v : want %v\n",
				test.description, test.peaks, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

func TestSolution(t *testing.T) {

	for k, test := range peaks[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solution(test.A); got != test.flag {

			t.Fatalf("FAIL: %s \n n: %v got %v : want %v\n",
				test.description, test.A, got, test.flag)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

func TestThirdpartySolution(t *testing.T) {

	for k, test := range peaks[:1] {

		fmt.Printf("k %v | test %v\n", k, test.description)

		if flag := thirdpartySolution(test.A); flag != test.flag {

			t.Fatalf("FAIL: %s \n A: %v got %v : want %v\n",
				test.description, test.A, flag, test.flag)
		}

	}
}
