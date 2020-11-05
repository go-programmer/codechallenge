package lesson

var sieve = []struct {
	description string
	T           []int
	want       	string
	err         string
}{
	{
		description: "Case 1",
		T:           []int{-3, -14, -5, 7, 8, 42, 8, 3},
		want:        "SUMMER",
		err:         "",
	},
	{
		description: "Case 2",
		T:           []int{2, -3, 3, 1, 10, 8, 2, 5, 13, -5, 3, -18},
		want:        "AUTUMN",
		err:         "",
	},
	
}
