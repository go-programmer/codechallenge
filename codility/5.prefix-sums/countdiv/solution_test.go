package countdiv

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) Test(c *C) {
	c.Assert(solution(1, 1000, 11), Equals, 90)
	c.Assert(solution(0, 0, 11), Equals, 1)

}

func (s *MySuite) Benchmark(c *C) {
	for i := 0; i < c.N; i++ {
		solution(1, 1000, 11)
	}
}
