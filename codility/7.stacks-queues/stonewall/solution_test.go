package stonewall

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestStonewall(c *C) {
	c.Assert(solution([]int{8, 8, 5, 7, 9, 8, 7, 4, 8}), Equals, 7)
}

func (s *MySuite) BenchmarkStonewall(c *C) {
	for i := 0; i < c.N; i++ {
		solution([]int{8, 8, 5, 7, 9, 8, 7, 4, 8})
	}
}
