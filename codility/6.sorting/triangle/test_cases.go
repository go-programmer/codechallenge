package triangle

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Example 1",
		A:           []int{10, 2, 5, 1, 8, 20},
		want:        1,
		err:         "",
	},
	{
		description: "Example 2",
		A:           []int{10, 50, 5, 1},
		want:        0,
		err:         "",
	},
}
