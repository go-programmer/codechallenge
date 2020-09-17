package uniquenumbers

var testCases = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "5 unique numbers from 1,2,3,4,5",
		A:           []int{1, 2, 2, 3, 4, 5},
		want:        5,
		err:         "",
	},
	{
		description: "3 unique numbers from 7,8,9",
		A:           []int{7, 8, 7, 9, 8, 7, 8, 9},
		want:        3,
		err:         "",
	},
}
