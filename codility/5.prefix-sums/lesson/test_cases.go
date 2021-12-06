package lesson

var testCases = []struct {
	description string
	A           []int
	k           int // Starting positino
	m           int // Moves that can be made
	want        int
	err         string
}{
	{
		description: "Example 1",
		A:           []int{2, 3, 7, 5, 1, 3, 9},
		k:           4,
		m:           6,
		want:        10,
		err:         "",
	},
}
