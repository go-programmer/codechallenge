package discintersections

import (
	"fmt"
	"sort"
)

/**
Complexity:
Expected worst-case time complexity is O(N*log(N));
Expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
*/

// Map circles to points on the left and on the right
// of the line segment -x, x and sort them. Then for
// every left point do a binary search for the index
// of the other point. Count the number of the
// elements before that index and subtract i+1 to
// remove duplicate counts.

func numberOfDiscIntersections(A []int) int {
	lenA := len(A)
	lefts := make([]int, lenA)
	rights := make([]int, lenA)
	result := 0

	for i := 0; i < lenA; i++ {
		lefts[i] = i - A[i]
		rights[i] = i + A[i]
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	for i := 0; i < lenA; i++ {
		// end := rights[i]

		// Binary search for index of element of
		// the right most value less than to the
		// interval end
		count := sort.Search(lenA, func(j int) bool {
			fmt.Printf("left %d, end %d\n", lefts[j], rights[i])
			return lefts[j] > rights[i]
		})

		count -= i + 1
		result += count

		if result > 10000000 {
			return -1
		}

	}

	return result
}

//package solution

//// you can also use imports, for example:
//// import "fmt"
//// import "os"
//
//import "sort"
//
//// you can write to stdout for debugging purposes, e.g.
//// fmt.Println("this is a debug message")
//
//func Solution(A []int) int {
//	pairs := [][]int{}
//	starts := make([]int, len(A))
//	result := 0
//
//	for i := 0; i < len(A); i += 1 {
//		pairs = append(pairs, []int{i - A[i], i + A[i]})
//	}
//
//	sort.Slice(pairs[:], func(i, j int) bool {
//		for x := range pairs[i] {
//			if pairs[i][x] == pairs[j][x] {
//				continue
//			}
//			return pairs[i][x] < pairs[j][x]
//		}
//		return false
//	})
//
//	for i := 0; i < len(A); i += 1 {
//		starts[i] = pairs[i][0]
//	}
//
//	for i := 0; i < len(A); i += 1 {
//		end := pairs[i][1]
//
//		// Binary search for index of element of the rightmost value less than to the interval-end
//		count := sort.Search(len(starts), func(i int) bool {
//			return starts[i] > end
//		})
//
//		count -= (i + 1)
//		result += count
//		if result > 10000000 {
//			return -1
//		}
//
//	}
//
//	return result
//}
