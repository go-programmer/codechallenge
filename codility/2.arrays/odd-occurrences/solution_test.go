package oddoccurences

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestSingleOddInteger(c *C) {
	testData := []int{5}
	want := 5
	got := solution(testData)

	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestOddAtMiddle(c *C) {
	testData := []int{1, 2, 3, 1, 2, 3, 7, 4, 5, 6, 4, 5, 6}
	want := 7
	got := solution(testData)

	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestFirstTwoOdds(c *C) {
	testData := []int{1, 2}
	want := 1
	got := solution(testData)

	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestSingleOddIntegerMapImpl(c *C) {
	testData := []int{5}
	want := 5
	got := solutionWithHashMap(testData)

	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestOddAtMiddleMapImpl(c *C) {
	testData := []int{1, 2, 3, 1, 2, 3, 7, 4, 5, 6, 4, 5, 6}
	want := 7
	got := solutionWithHashMap(testData)

	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestFirstTwoOddsMapImpl(c *C) {
	testData := []int{1, 2}
	want := 1
	got := solutionWithHashMap(testData)

	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestOddNumber(c *C) {
	testData := 1
	want := true
	got := isOdd(testData)

	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestEvenNumber(c *C) {
	testData := 2
	want := false
	got := isOdd(testData)

	c.Assert(got, Equals, want)
}

func (s *TestSuite) BenchmarkOddOccurences(c *C) {
	for i := 0; i < c.N; i++ {
		solutionWithHashMap([]int{1, 1, 5, 2, 2, 3, 5, 7, 3})
	}
}
