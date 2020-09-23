package lesson

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Example 1",
		A:           []int{5, -7, 3, 5, -2, 4, -1},
		want:        10,
		err:         "",
	},
}
