package fish

var testCases = []struct {
	description string
	A           []int
	B           []int
	want        int
	err         string
}{
	{
		description: "Given example",
		A:           []int{4, 3, 2, 1, 5},
		B:           []int{0, 1, 0, 1, 0},
		want:        2,
		err:         "",
	},
	{
		description: "All fish up",
		A:           []int{4, 3, 2, 1, 5},
		B:           []int{0, 0, 0, 0, 0},
		want:        5,
		err:         "",
	},
	{
		description: "All fish down",
		A:           []int{4, 3, 2, 1, 5},
		B:           []int{1, 1, 1, 1, 1},
		want:        5,
		err:         "",
	},
	{
		description: "First up and down",
		A:           []int{4, 3, 2, 1, 5},
		B:           []int{1, 0, 0, 1, 0},
		want:        2,
		err:         "",
	},
	{
		description: "5 kill 1 and 4",
		A:           []int{4, 3, 2, 1, 5, 6, 9},
		B:           []int{1, 0, 0, 1, 0, 1, 0},
		want:        2,
		err:         "",
	},
}
