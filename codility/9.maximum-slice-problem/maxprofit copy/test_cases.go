package solution

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Example provided",
		A:           []int{23171, 21011, 21123, 21366, 21013, 21367},
		want:        356,
		err:         "",
	},

	{
		description: "Loss",
		A:           []int{6, 5, 1},
		want:        0,
		err:         "",
	},
}
