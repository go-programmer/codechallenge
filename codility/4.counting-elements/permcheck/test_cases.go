package permcheck

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Example test case returns 1.",
		A:           []int{4, 1, 2, 3},
		want:        1,
		err:         "",
	},
	{
		description: "Example test case returns 0.",
		A:           []int{4, 1, 3},
		want:        0,
		err:         "",
	},
}
