package cyclicrotation

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestCyclicRotation(c *C) {
	c.Assert(solution([]int{0 - 9}, 2), DeepEquals, []int{-9})
	c.Assert(solution([]int{0 - 9}, 1), DeepEquals, []int{-9})
	c.Assert(solution([]int{1, 1, 2, 3, 5}, 42), DeepEquals, []int{3, 5, 1, 1, 2})
	c.Assert(solution([]int{3, 8, 9, 7, 6}, 3), DeepEquals, []int{9, 7, 6, 3, 8})
}

func (s *TestSuite) BenchmarkCyclicRotation(c *C) {
	for i := 0; i < c.N; i++ {
		solution([]int{3, 8, 9, 7, 6}, 3)
	}
}

// Resources:
// https://leetcode.com/problems/rotate-array/discuss/806176/20ms-in-C-with-Time-O(n)-space-O(1)-accepted-solution
// https://www.geeksforgeeks.org/reversal-algorithm-right-rotation-array
// https://www.techiedelight.com/right-rotate-an-array-k-times
// https://solution.programmingoneonone.com/2020/07/hackerrank-arrays-left-rotation-solution.html
// https://www.javatpoint.com/java-program-to-right-rotate-the-elements-of-an-array
// https://www.geeksforgeeks.org/python-program-for-program-for-array-rotation-2
// https://www.geeksforgeeks.org/reversal-algorithm-right-rotation-array
// https://stackoverflow.com/questions/21322173/convert-rune-to-int
