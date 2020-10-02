package solution

var arrayA = []struct {
	description string
	A           []int
	want        []int
	err         string
}{
	{
		description: "Example",
		A:           []int{3, 1, 2, 3, 6},
		want:        []int{2, 4, 3, 2, 0},
		err:         "",
	},

	{
		description: "All same elements",
		A:           []int{3, 3, 3, 3, 3},
		want:        []int{0, 0, 0, 0, 0},
		err:         "",
	},

	{
		description: "All prime numbers",
		A:           []int{2, 3, 5, 7},
		want:        []int{3, 3, 3, 3},
		err:         "",
	},
}
