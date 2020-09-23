package equileader

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Example 1",
		A:           []int{4, 3, 4, 4, 4, 2},
		want:        2,
		err:         "",
	},
}
