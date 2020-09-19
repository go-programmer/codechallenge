package minavgtwoslice

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) Test(c *C) {
	A := []int{4, 2, 2, 5, 1, 5, 8}
	c.Assert(solution(A), Equals, 1)
}

func TestAvg(mt *testing.T) { TestingT(mt) }

type MinAvgSuite struct{}

var _ = Suite(&MinAvgSuite{})

func (s *MinAvgSuite) TestingT(c *C) {
	A := []int{4, 2, 2, 5, 1, 5, 8}
	P := prefixSum(A)

	c.Assert(getMinAverage(P[2:], P[0]), Equals, float64(2))
}

// func (s *MySuite) Benchmark(c *C) {
// 	for i := 0; i < c.N; i++ {
// 		solution(1, 1000, 11)
// 	}
// }
