package passingcars

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestFrogRiverOne(c *C) {
	c.Assert(solution([]int{0, 1, 0, 1, 1}), Equals, 5)
}

func (s *MySuite) BenchmarkFrogRiverOne(c *C) {
	for i := 0; i < c.N; i++ {
		solution([]int{0, 1, 0, 1, 1})
	}
}
