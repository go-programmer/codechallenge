package solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	// len(blocks)-1
	for k, test := range blocks[:] {

		fmt.Printf("k %v | test %v\n", k, test.description)
		blocks := solution(test.A)

		if blocks != test.blocks {

			t.Fatalf("FAIL: %s \n A: %v got %v : want %v\n",
				test.description, test.A, blocks, test.blocks)
		}

	}
}

func TestPeaks(t *testing.T) {

	for k, test := range peaks[len(peaks)-3:] {

		fmt.Printf("k %v | test %v\n", k, test.description)
		peaks := countPeaks(test.A, len(test.A))

		if !reflect.DeepEqual(peaks, test.peaks) {

			t.Fatalf("FAIL: %s \n A: %v got %v : want %v\n",
				test.description, test.A, peaks, test.peaks)
		}

	}
}

func TestFactorials(t *testing.T) {

	for k, test := range factors[:] {

		fmt.Printf("k %v | test %v\n", k, test.description)
		factorials, _ := factorials(test.A)

		if !reflect.DeepEqual(factorials, test.factorials) {

			t.Fatalf("FAIL: %s \n A: %v got %v : want %v\n",
				test.description, test.A, factorials, test.factorials)
		}

	}
}
