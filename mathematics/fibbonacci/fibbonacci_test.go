package fibbonacci

import (
	"reflect"
	"testing"
)

func TestFibbonacciSerial(t *testing.T) {
	want := []int{1, 1, 2}
	got := fibbonacci(3)
	gotLen := len(got)
	wantLen := len(want)

	if gotLen != wantLen {
		t.Errorf("TestFibbonacciSerial len(got) %v != len(want) %v", gotLen, wantLen)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestFibbonacciSerial got %v, want %v", got, want)
	}

	for i, fib := range got {

		if fib != want[i] {
			t.Errorf("TestFibbonacciSerial got %v, want %v", fib, want)
		}
	}

}
