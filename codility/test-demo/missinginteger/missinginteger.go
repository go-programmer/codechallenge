package missinginteger

import (
	"fmt"
	"sort"
)

// Problem:
// Write a function:

// func Solution(A []int) int

// that, given an array A of N integers, returns
// the smallest positive integer (greater than 0)
// that does not occur in A.
// For example,
// Given A = [1, 3, 6, 4, 1, 2], the function should return 5.
// Given A = [1, 2, 3], the function should return 4.
// Given A = [−1, −3], the function should return 1.
// Write an efficient algorithm for the following assumptions:
// N is an integer within the range [1..100,000];
// Each element of array A is an integer within the range [−1,000,000..1,000,000].

// missingInteger Find the smallest positive integer
// (greater than 0) that does not occur in a slice of int
// Score 55%
func missingInteger(A []int) int {
	// write your code in Go 1.4

	sort.Ints(A)
	minVal := 1

	if len(A) == 1 && (A[0] >= 1000000 || A[0] < 0) {
		return 1
	}

	if A[len(A)-1] <= 0 {
		return 1
	}

	lenA := len(A) - 2
	i := 0
	for i <= lenA {
		if (A[i] >= 0) && (A[i+1]-A[i] > 1) {
			return A[i] + 1
		}
		i++
	}

	if (i-1 == lenA) && (A[i] > minVal) {
		return A[i] + 1
	}

	return minVal
}

// Client calls missingInteger
func Client() {
	caseOne()
	caseTwo()
	caseThree()
	caseFour()
	caseFive()
}

func caseFive() {
	a := []int{-1, -2000, -3}
	ret := missingInteger(a)

	if ret == 1 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
func caseFour() {
	a := []int{-1, 000, 000}
	ret := missingInteger(a)

	if ret == 1 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
func caseThree() {

	a := []int{1, 3, 6, 4, 1, 2, 10000000}
	ret := missingInteger(a)

	if ret == 5 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
func caseTwo() {
	a := []int{9, 3, 6, 1, 15, 21, 8, 36, 45, 5}
	ret := missingInteger(a)

	if ret == 2 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
func caseOne() {
	a := []int{-1, -50, 1, 2, 3}
	ret := missingInteger(a)

	if ret == 4 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
