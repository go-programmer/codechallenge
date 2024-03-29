package stockslaundering

import (
	"fmt"
	"sort"
)

// Problem:
// Write a function:
//     class Solution { public int solution(int K, int[] C, int[] D); }

// that, given an integer K (the number of socks that the washing machine can clean),
// two arrays C and D (containing the color representations of N clean and M dirty socks
// respectively), returns the maximum number of pairs of socks.
// For example, given K = 2, C = [1, 2, 1, 1] and D = [1, 4, 3, 2, 4],
// the function should return 3, as explained above.

// Assume that:
//         K is an integer within the range [0..50];
//         each element of arrays C, D is an integer within the range [1..50];
//         C and D are not empty and each of them contains at most 50 elements.

// https://app.codility.com/programmers/lessons/92-tasks_from_indeed_prime_2016_college_coders_challenge/socks_laundering/

// Solution scored 70%. It failed on:
// there are no clean socks taken: got 3 expected 2
// randomly generated tests with only a couple of colors: got 9 expected 8
// maximal possible test cases: got 36 expected 37; got 36 expected 33

func stocksLaundering(K int, C []int, D []int) int {
	lenC := len(C)
	lenD := len(D)

	if lenC == lenD && lenC < 2 && C[0] != D[0] {
		return 0
	}

	sort.Ints(C)
	var singleCS []int
	totalPair := 0

	// fmt.Println()
	// fmt.Printf("C:%v len(C):%v\n", C, lenC)

	if lenC > 1 {
		for i := 0; i < lenC-1; i++ {

			if C[i] == C[i+1] {
				totalPair++
				i = i + 1

			} else {
				singleCS = append(singleCS, C[i])
			}
		}

		if C[lenC-2] != C[lenC-1] {
			singleCS = append(singleCS, C[lenC-1])
		}
	} else {
		singleCS = append(singleCS, C[0])
	}

	sort.Ints(singleCS)
	// fmt.Printf("Total pair %v\n", totalPair)
	// fmt.Printf("Single CS %v\n", singleCS)
	// fmt.Printf("K:%v\n", K)

	if K == 0 {
		return totalPair
	}

	sort.Ints(D)
	dsMap := make(map[int]int)
	for i := 0; i < lenD; i++ {
		if _, ok := dsMap[D[i]]; !ok {
			dsMap[D[i]] = 1
		} else {
			dsMap[D[i]]++
		}
	}

	for _, scs := range singleCS {

		if _, ok := dsMap[scs]; ok {
			totalPair++
			K--
			dsMap[scs]--

			if dsMap[scs] == 0 {
				delete(dsMap, scs)

				if len(dsMap) == 0 {
					return totalPair
				}
			}

			if K == 0 {
				return totalPair
			}
		}
	}

	// fmt.Printf("DMap %v\n", dsMap)
	// fmt.Printf("Total %v\n", totalPair)
	for _, v := range dsMap {
		// fmt.Printf("v %v\n", v)

		if v >= 2 && K >= 2 {
			divisor := v / 2
			mod := v % 2
			// fmt.Printf("div %v\n", divisor)
			if K >= divisor {
				totalPair = totalPair + divisor

				if mod == 0 {
					K -= v
				} else {
					K -= v - 1
				}

				if K < 2 {
					return totalPair
				}
			}
		}
	}

	// fmt.Printf("DMap after cs  %v\n", dsMap)
	return totalPair
}

// PrintFail Print the failed case
func PrintFail(got int, want int, failMsg string) {
	if got != want {
		fmt.Printf("FAIL %v g:%d w:%d\n", failMsg, got, want)
	}
}

// THOUGHT PROCESS

//THOUGHT 2 : Use map
//COUNT CLEAN SOCKS PAIR
// Sort clean socks.
// Count total number of pair of clean socks.
// Store the single clean socks numbers in a slice.

//COUNT CLEAN AND DIRTY SOCKS PAIR
// Note: The number of dirty socks being paired to clean socks
// should be less than or equal to machine limit!
// Sort single clean socks slice.
// Sort ds.

// If dirty socks min > clean socks max there can not be any pair!
// Else, find the match.

// Loop ds
// Store ds in a map,
//  The key of the map is socks
//  The value is total number of same socks

// Loop single clean socks slice.
//  Compare with dirty socks.
//  If any match is found,
// 	 Increase total clean socks pair.
//   Decrease the machine limit.
// 	 Decrease the ds map value by 1.
//   If ds number == 0, remove that ds
//   If the machine limit is reached ml=0 or 1, return clean socks pair.

//COUNT DIRTY SOCKS PAIR
// We need to find the dirty socks based on the remaining machine limit!
// Range ds map.
// If value is greater than or equal to 2// we only need pairs
//  Divide the value by two, store the divisor
// 	If ml is greater than divisor,
//		Add divisor to the total clean socks pair.
// 		Decrease ml, substract divisor.
// 		No need to update the ds value as we will move to next ds
// If ml < 2 return return clean socks pair.
//

//THOUGHT 1: Has O(N^2) behavior so skipped implementation
//COUNT CLEAN SOCKS PAIR
// Sort clean socks.
// Count total number of pair of clean socks.
// Store the single clean socks numbers in a slice.

//COUNT CLEAN AND DIRTY SOCKS PAIR
// The number of dirty socks being paired to clean socks
// should be less than or equal to machine limit!
// Sort dirty socks.
// If dirty socks min > clean socks max there can not be any pair!
// Else, find the match.
// Loop single clean socks slice.
//  Compare with dirty socks.
//  If any match is found,
// 	 Increase total clean socks pair,
//   Decrease the machine limit.
//   Set the value to 0 to the dirty sock.
// 	 Skip the same value. Increase the index. ** THIS WILL BE A ANOTHER LOOP! **
//  If the machine limit is reached, return clean socks pair.
//  If dirty socks slice end is reached, break the loop.

//COUNT DIRTY SOCKS PAIR
// We need to find the dirty socks based on the remaining machine limit!
// While machine limit is greater than 0,
// If pair found, (current == next)
//  Increase total clean socks pair
//  Decrease machine limit by two.
// 	Incrment i+1
//  If dirty socks slice end is reached
// 		Return total clean socks pair.

// THOUGHT 0
// if cs pair found ( index i == next index i+1 ):
// 		increase tcs(total clean socks);
// 		move the index next of next index  ( i = next + 1 )
// else ( the cs is single )
// 		check if a pick dirty socks that can be paired with it
// 		if pair found (true):
// 			increment tws by 1
// 			mark the index of that dirty sock,
//
//HAD another concept now. This was getting complicated.
//
// Check if there is a dirty socks pair for those single clean socks; count them  say singleDirtySocks
// Count total number of socks
// Pick the pair of remaining dirty socks, say dirtySocksPair
// Total dirty socks must be <= The washing machine's limit
