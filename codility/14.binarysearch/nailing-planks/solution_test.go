package nailingplanks

import (
	"fmt"
	"testing"
)

// func TestNailed(t *testing.T) {
// 	for k, test := range nailedCases[:] {
// 		fmt.Printf("k %v | Test %v\n", k, test.description)

// 		if got := getNailed(test.low, test.high, test.nails); got != test.want {
// 			t.Fatalf("FAIL: %s \n Low %#v\n High %#v\n Nails %#v \n got %d : want %d\n",
// 				test.description, test.low, test.high, test.nails, got, test.want)
// 		}
// 		t.Logf("SUCCESS: %s", test.description)
// 	}

// }

func TestNailingPlanks(t *testing.T) {
	for k, test := range testCases {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := thirdSolution(test.plankA, test.plankB, test.nails); got != test.want {
			t.Fatalf("FAIL: %s \n PlankA %#v\n PlankB %#v\n Nails %#v \n got %d : want %d\n",
				test.description, test.plankA, test.plankB, test.nails, got, test.want)
		}
		t.Logf("SUCCESS: %s", test.description)
	}
}
