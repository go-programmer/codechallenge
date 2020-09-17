package tapeequlibrium

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestSolution(c *C) {
	c.Assert(solution([]int{3, 1, 2, 4, 3}), Equals, 1)
}

func (s *MySuite) BenchmarkSolution(c *C) {
	for i := 0; i < c.N; i++ {
		solution([]int{3, 1, 2, 4, 3})
	}
}
