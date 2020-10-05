package solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	for k, test := range semiprime[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := Solution(test.N, test.P, test.Q); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("FAIL: %s \n N: %#v got %v : want %v\n",
				test.description, test.N, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

// func TestSmallestPrimeFactor(t *testing.T) {
// 	for k, test := range smallestPrimeFactor[1:] {
// 		fmt.Printf("k %v | test %v\n", k, test.description)

// 		if got := smallestPrimeFactors(test.n); !reflect.DeepEqual(got, test.want) {
// 			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
// 				test.description, test.n, got, test.want)
// 		}

// 		// t.Logf("SUCCESS: %s", test.description)
// 	}
// }

// func TestPrimeFactor(t *testing.T) {
// 	for k, test := range primeFactor[:] {
// 		fmt.Printf("k %v | test %v\n", k, test.description)

// 		if got := primeFactors(test.n); !reflect.DeepEqual(got, test.want) {
// 			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
// 				test.description, test.n, got, test.want)
// 		}

// 		// t.Logf("SUCCESS: %s", test.description)
// 	}
// }
