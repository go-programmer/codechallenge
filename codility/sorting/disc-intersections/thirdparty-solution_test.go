package discintersections

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
	c.Assert(numberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0}), Equals, 11)
}

func (s *MySuite) BenchmarkNumberOfDiscIntersections(c *C) {
	for i := 0; i < c.N; i++ {
		numberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0})
	}
}
