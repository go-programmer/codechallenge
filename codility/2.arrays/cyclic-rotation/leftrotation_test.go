package cyclicrotation

import (
	"fmt"
	"testing"
	"reflect"
)

// 9 has binary representation 1001 and contains a binary gap of length 2

var testcase = []struct {
	description string
	A           []int
	K			int	
	want        []int
	err         string
}{
	{
		description: "Provided example",
		A:           []int{3, 8, 9, 7, 6},
		K:			 1,
		want:        []int{8,9,7,6,3},
		err:         "",
	},
	{
		description: "Provided example",
		A:           []int{3, 8, 9, 7, 6},
		K:			 2,
		want:        []int{9,7,6,3,8},
		err:         "",
	},
	{
		description: "Provided example",
		A:           []int{3, 8, 9, 7, 6},
		K:			 3,
		want:        []int{7,6,3,8,9},
		err:         "",
	},
	{
		description: "Provided example",
		A:           []int{3, 8, 9, 7, 6},
		K:			 10,
		want:        []int{3,8,9,7,6},
		err:         "",
	},
	{
		description: "Provided example",
		A:           []int{3, 8, 9, 7, 6},
		K:			 6,
		want:        []int{8,9,7,6,3},
		err:         "",
	},
	{
		description: "Provided example",
		A:           []int{1,2,3,4},
		K:			 1,
		want:        []int{2,3,4,1},
		err:         "",
	},
}

func TestSolution(t *testing.T) {
	for k, test := range testcase[:] {
		fmt.Printf("k %v | test %v\n", k, test.description)

		if got := leftRotation(test.A, test.K); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("FAIL: %s \n n: %#v got %v : want %v\n",
				test.description, test.A, got, test.want)
		}

		// t.Logf("SUCCESS: %s", test.description)
	}
}
