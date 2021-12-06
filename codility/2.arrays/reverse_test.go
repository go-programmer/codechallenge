package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func reverse(A []int) []int {
	lenA := len(A)

	for i := 0; i < lenA/2; i++ {
		k := lenA - i - 1
		A[i], A[k] = A[k], A[i]
	}

	for _, v := range A {
		fmt.Println(v)
	}

	return A
}

func TestReverse(t *testing.T) {
	want := []int{5, 3, 2, 1}
	if got := reverse([]int{1, 2, 3, 5}); !reflect.DeepEqual(got, want) {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}
}
