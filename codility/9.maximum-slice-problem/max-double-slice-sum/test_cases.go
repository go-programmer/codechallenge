package solution

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Example provided",
		A:           []int{3, 2, 6, -1, 4, 5, -1, 2},
		want:        17,
		err:         "",
	},
}
