package maxcounter

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaxCounters(c *C) {
	c.Assert(solution(5, []int{3, 4, 4, 6, 1, 4, 4}), DeepEquals, []int{3, 2, 2, 4, 2})
}

// func (s *MySuite) BenchmarkMaxCounters(c *C) {
// 	for i := 0; i < c.N; i++ {
// 		thirdpartySolution(5, []int{1, 3, 6, 4, 1, 2})
// 	}
// }
