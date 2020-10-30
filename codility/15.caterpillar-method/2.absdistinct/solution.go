package solution

import (
	"math"
)

func solution(A []int) int {
	set := NewSet()	

	for _, a := range A {
		absVal:= math.Abs(float64(a))

		if !set.Contains(absVal) {
			set.Add(absVal)
		}
	}

	return len(set.m)
}


// Set is a custom `set` datatype
type Set struct {
    m map[float64]struct{}
}

var void = struct{}{}

// NewSet returns pointer to set
func NewSet() *Set {
    s := &Set{}
    s.m = make(map[float64]struct{})
    return s
}

// Add value to s
func (s *Set) Add(value float64) {
    s.m[value] = void
}

// Remove value
func (s *Set) Remove(value float64) {
    delete(s.m, value)
}

//Contains returns true if value exists in s
func (s *Set) Contains(value float64) bool {
    _, c := s.m[value]
    return c
}