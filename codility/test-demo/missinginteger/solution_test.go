package missinginteger

import (
	"testing"
)

func TestShouldReturnOne(t *testing.T) {
	a := []int{-1, -1, -1, -1, -1, 2, 3, 4, 4}
	want := 1

	if got := solution(a); got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}

}

func TestSecondLastAndLastValueAreEqual(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 2, 3, 4, 4, 4}
	want := 5

	if got := solution(a); got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}

}

func TestSingleElement(t *testing.T) {
	a := []int{1}
	want := 2

	if got := solution(a); got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}

}

func TestCaseSix(t *testing.T) {
	a := []int{1}
	got := solution(a)
	want := 2
	if got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}
}

func TestCaseFive(t *testing.T) {
	a := []int{-1, -2000, -3}
	got := solution(a)
	want := 1
	if got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}
}

func TestCaseFour(t *testing.T) {
	a := []int{-1, 000, 000}
	got := solution(a)
	want := 1

	if got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}
}

func TestCaseThree(t *testing.T) {
	a := []int{1, 3, 6, 4, 1, 2, 10000000}
	got := solution(a)
	want := 5

	if got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}
}

func TestCaseTwo(t *testing.T) {
	a := []int{9, 3, 6, 1, 15, 21, 8, 36, 45, 5}
	got := solution(a)
	want := 2

	if got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}
}
func TestCaseOne(t *testing.T) {
	a := []int{-1, -50, 1, 2, 3}
	got := solution(a)
	want := 4

	if got != want {
		t.Fatalf("FAIL:  got %d != want %d\n", got, want)
	}
}
