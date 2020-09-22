package dominator

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Example 1",
		A:           []int{3, 4, 3, 2, 3, -1, 3, 3},
		want:        7,
		err:         "",
	},
}
