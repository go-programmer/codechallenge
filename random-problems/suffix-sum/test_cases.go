package suffixsum

var testCases = []struct {
	description string
	array       []int
	want        []int
	err         string
}{

	{
		description: "Example 1",
		array:       []int{1, 2, 3, 4},
		want:        []int{10, 9, 7, 4},
		err:         "",
	},
}
