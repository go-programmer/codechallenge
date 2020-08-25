package frogriverone

import (
	"fmt"
	"sort"
)

// Problem:
// You are given an array A consisting of N integers
// representing the falling leaves. A[K] represents the
// position where one leaf falls at time K, measured in seconds.

// >> The frog can cross only when leaves appear
// >> at every position across the river from 1 to X.

// X is the position of the frog. Find K.

// Attempts made 25 Aug 2020
// First attempt [Score 18%] https://app.codility.com/demo/results/trainingRYAGW8-VUP
// Second attempt [Score 18%] https://app.codility.com/demo/results/trainingKHUKRM-EJ2
// Third attempt [Score 27%] https://app.codility.com/demo/results/training9T3V7Y-3PB
// Fourth attempt [Score 27%] https://app.codility.com/demo/results/trainingWCYDG3-EV3
// Fifth attempt [Score 54%] https://app.codility.com/demo/results/trainingU5BSRF-2E7
// Sixth attempt [Score 81%] https://app.codility.com/demo/results/trainingGHNZRQ-EKW
// Seventh attempt [Score 81%] https://app.codility.com/demo/results/trainingFFTAHH-Y7Q
// Eight attempt [Score 81%] https://app.codility.com/demo/results/trainingTXWBEQ-FKU
// 100% https://app.codility.com/demo/results/trainingXW3A5E-KR2

func frogRiverOneClient() {
	var testcase string
	var want, got int
	var A []int

	testcase = "1"
	A = []int{1, 3, 1, 4, 2, 4, 5, 4}
	want = 6
	got = frogRiverOne(5, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

	testcase = "2" // Leaf falling at where the frog is
	A = []int{2, 2, 2, 2, 2}
	want = -1
	got = frogRiverOne(2, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

	testcase = "3"
	A = []int{1, 2}
	want = 1
	got = frogRiverOne(2, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

	testcase = "4"
	A = []int{5, 1, 2, 1, 4, 2, 4, 5, 3}
	want = 8
	got = frogRiverOne(3, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

	testcase = "5 Sequential"
	A = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 11}
	want = 11
	got = frogRiverOne(11, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

	// in this case, the remaining steps were found after
	// the position the frog was at.
	testcase = "6"
	A = []int{1, 3, 1, 3, 2, 1, 3}
	want = 4
	got = frogRiverOne(3, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

	testcase = "7 Single" // 1 element
	A = []int{5}
	want = 0
	got = frogRiverOne(5, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

	testcase = "8 Frog never crosses the river"
	A = []int{1, 3, 1, 10, 6, 5, 10, 21, 4, 5, 9, 6, 7}
	want = -1
	got = frogRiverOne(10, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

	testcase = "9 Wrong single element "
	A = []int{1}
	want = -1
	got = frogRiverOne(2, A)
	fmt.Printf("Test case:%v g:%d w:%d %v \n", testcase, got, want, got == want)

}

// Algorithm
// copy the value of A to a new array to preserve the index order
// sort A
//
// check continuous sequence exists to X
// while checking create a map: This might not be required!
//
// if sequence is continuous:
// 		all the values must occur
//
// add to map
// return the highest value of key in the map

// Solution scores 100% in Codility Arrays CyclicRotation
// Detected time complexity: O(N)

func frogRiverOne(X int, A []int) int {
	maxIndex := -1

	var magicArr [100000]int
	lenA := len(A)

	if lenA == 1 && A[0] == X {
		return 0
	}

	if lenA == 1 && A[0] != X {
		return -1
	}

	copy(magicArr[:], A)
	cpA := magicArr[:lenA]
	sort.Ints(cpA)

	if cpA[0] > 1 || cpA[0] == cpA[lenA-1] {
		return maxIndex
	}

	stepMap := make(map[int]int)
	sumOfStepMapKeys := 0

	for kA, vA := range A {

		if _, ok := stepMap[vA]; !ok && vA <= X {
			stepMap[vA] = kA

			if maxIndex < kA {
				maxIndex = kA
			}
			sumOfStepMapKeys += vA

			// fmt.Printf("k:sum %v:%v\n", v, sumOfStepMapKeys)
		}
	}

	// fmt.Println()

	// fmt.Printf("X %v Map %v\n", X, stepMap)

	sumOfSequentialNum := (X * (X + 1)) / 2
	// fmt.Printf("Sum Seq %d : Map %d\n", sumOfSequentialNum, sumOfStepMapKeys)
	if sumOfSequentialNum == sumOfStepMapKeys {
		return maxIndex
	}

	return -1
}

// The solution passes all example tests
// but fails the performance times out with error
// Detected time complexity: O(N ** 2)

func frogRiverOnePerformaceFail(X int, A []int) int {

	// make a key value pair KVP as key is the value of A, and value is the index
	// sort A
	// search X; if X exists return value from KVP(X)
	// In the above algo; we have to traverse through all the value while making KVP.
	// So does not sound efficient.

	// Linear O(N) bruteforce solution
	// Search for X in A

	// Aux map to check if lower steps were exists
	lowerSteps := make(map[int]bool)
	step := 1
	for k, v := range A {

		// Previous logic that fails when lower steps comes after X, Case 6
		// if lower step exists dont enter in map

		// check if map exists in the map
		// add to map if step does not exist
		// after addging; check if the all the steps lower than X exists
		if _, ok := lowerSteps[v]; !ok && v <= X {
			lowerSteps[v] = true

			for i := 0; i < X; i++ {

				if _, ok := lowerSteps[step]; ok {

					if step == X {
						return k
					}
					step++
				}
			}
		}

	}
	return -1
}

// Intial attempt
// Fails correctness and performance test cases
// Time complexity O(N ** 2)
func frogRiverOnePerInitial(X int, A []int) int {

	// make a key value pair KVP as key is the value of A, and value is the index
	// sort A
	// search X; if X exists return value from KVP(X)
	// In the above algo; we have to traverse through all the value while making KVP.
	// So does not sound efficient.

	// Linear O(N) bruteforce solution
	// Search for X in A

	// Aux map to check if lower steps were exists
	lowerSteps := make(map[int]bool)
	step := 1
	for k, v := range A {

		// Previous logic that fails when lower steps comes after X, Case 6
		// if lower step exists dont enter in map
		if _, ok := lowerSteps[v]; !ok && v <= X {
			lowerSteps[v] = true
		}

		if X == v {
			// check if the all the lower steps exists lower than X
			for step < X {

				if _, ok := lowerSteps[step]; !ok {
					break
				}
				step++
			}

			// All steps exists
			if step == X {
				return k
			}
		}

		// if X exists in the map already,
		// check if all lower steps exists
		if _, ok := lowerSteps[X]; ok {

			for mk := range lowerSteps {

				if _, ok := lowerSteps[step]; ok {
					// if step == X; return k
					if step == X {
						return k
					}
					step++
				}
			}
		}
	}

	return -1
}
