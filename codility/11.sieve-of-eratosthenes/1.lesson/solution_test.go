package lesson

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSeieve(t *testing.T) {
	for k, test := range sieve[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := markPrime(test.n); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

func TestSmallestPrimeFactor(t *testing.T) {
	for k, test := range smallestPrimeFactor[1:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := smallestPrimeFactors(test.n); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}

func TestPrimeFactor(t *testing.T) {
	for k, test := range primeFactor[1:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := primeFactors(test.n); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.n, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
