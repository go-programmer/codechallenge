package solution

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Example lesson",
		A:           []int{5, -7, 3, 5, -2, 4, -1},
		want:        10,
		err:         "",
	},
	{
		description: "Example provided",
		A:           []int{3, 2, -6, 4, 0},
		want:        5,
		err:         "",
	},
	{
		description: "1 element",
		A:           []int{1},
		want:        1,
		err:         "",
	},
	{
		description: "-1 element",
		A:           []int{-1},
		want:        -1,
		err:         "",
	},

	{
		description: "Two -ve elements",
		A:           []int{-1, -1000},
		want:        -1,
		err:         "",
	},
}
