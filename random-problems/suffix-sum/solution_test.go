package suffixsum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	for k, test := range testCases {

		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := solutionSameSlice(test.array); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Error %s: got:want %v:%v\n", test.description, got, test.want)
		}
	}
}
