package discintersections

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Should return 11",
		A:           []int{1, 5, 2, 1, 4, 0},
		want:        11,
		err:         "",
	},
	// {
	// 	description: "3 unique numbers from 7,8,9",
	// 	A:           []int{7, 8, 7, 9, 8, 7, 8, 9},
	// 	want:        3,
	// 	err:         "",
	// },
}
