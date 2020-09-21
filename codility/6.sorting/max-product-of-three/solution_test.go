package distinct

import (
	"testing"

	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func Test(t *testing.T) {
	TestingT(t)
}

func (s *MySuite) TestNumberOfDiscIntersections(c *C) {

	// Test case 1
	// c.Assert(solution([]int{-3, 1, 2, -2, 5, 6}), Equals, 60)

	// Failed case while submission
	// c.Assert(solution([]int{-5, 5, -5, 4}), Equals, 125)
	c.Assert(solution([]int{-5, -6, -4, -7, -10}), Equals, -120)

}

// func (s *MySuite) BenchmarkNumberOfDiscIntersections(c *C) {
// 	for i := 0; i < c.N; i++ {
// 		numberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0})
// 	}
// }
