package cyclicrotation

import "fmt"

// Solution scores 100% in Codility Arrays CyclicRotation
// Time O(n), space O(1)
// Resources:
// https://leetcode.com/problems/rotate-array/discuss/806176/20ms-in-C-with-Time-O(n)-space-O(1)-accepted-solution
// https://www.geeksforgeeks.org/reversal-algorithm-right-rotation-array
// https://www.techiedelight.com/right-rotate-an-array-k-times
// https://solution.programmingoneonone.com/2020/07/hackerrank-arrays-left-rotation-solution.html
// https://www.javatpoint.com/java-program-to-right-rotate-the-elements-of-an-array
// https://www.geeksforgeeks.org/python-program-for-program-for-array-rotation-2
// https://www.geeksforgeeks.org/reversal-algorithm-right-rotation-array
// https://stackoverflow.com/questions/21322173/convert-rune-to-int
func cyclicRotation(nums []int, k int) []int {
	numsSize := len(nums)

	if k == 0 || k == numsSize || k > 100 || numsSize == 1 || numsSize == 0 || numsSize > 100 {
		return nums
	}

	if k > numsSize {

		mod := k % numsSize
		fmt.Println(mod)

		if mod == 0 {
			k = k / numsSize
		} else {
			k = mod
		}
	}

	fmt.Println(k)

	rev(nums, 0, numsSize-1)
	rev(nums, 0, k-1)
	rev(nums, k, numsSize-1)

	return nums
}

func rev(a []int, b, e int) {
	for b < e {
		t := a[b]
		a[b] = a[e]
		a[e] = t
		b++
		e--
	}
}

// rotate array as given k times
func cyclicRotationClient() {
	// A := []int{3, 8, 9}
	// K := 2

	// OUT BOUND EXCEPTION
	// Fixing k value fixed it
	// A := []int{3, 8, 9}
	// K := 7

	// OUT BOUND EXCEPTION
	A := []int{3, 8}
	K := 5

	cyclicRotation(A, K)
	fmt.Println(A)

	// brute force ; need more math on this
	// swap with next element swap(k, k+1)
	// if last element, swap with the first element

	// if the rotation == lenght of array, no need to rotate
	// if len == 1 return array

}
