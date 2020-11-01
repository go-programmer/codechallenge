package solution

func solution(M int, A []int) int {
	lenA := len(A)

	if lenA >= 1000000000 {
		return 1000000000
	}
	 
	count := 0

	for i:=0; i<lenA; i++ {

		count++
		set := NewSet()	

		for j:=i+1; j<lenA; j++ {
	
			if set.Contains(A[j]) || A[j] == A[i]{
				break
			}

			set.Add(A[j])
			count++

		}

	}

	return count
}



// Set is a custom `set` datatype
type Set struct {
    m map[int]struct{}
}

var void = struct{}{}

// NewSet returns pointer to set
func NewSet() *Set {
    s := &Set{}
    s.m = make(map[int]struct{})
    return s
}

// Add value to s
func (s *Set) Add(value int) {
    s.m[value] = void
}

// Remove value
func (s *Set) Remove(value int) {
    delete(s.m, value)
}

//Contains returns true if value exists in s
func (s *Set) Contains(value int) bool {
    _, c := s.m[value]
    return c
}