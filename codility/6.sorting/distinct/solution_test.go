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
	c.Assert(solution([]int{0, 1, 2, 3}), Equals, 4)
}

// func (s *MySuite) BenchmarkNumberOfDiscIntersections(c *C) {
// 	for i := 0; i < c.N; i++ {
// 		numberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0})
// 	}
// }
