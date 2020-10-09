package solution

var solution = []struct {
	description string
	K           int
	M           int
	A           []int
	want        int
	err         string
}{
	{
		description: "Example",
		K:           3,
		M:           5,
		A:           []int{2, 1, 5, 1, 2, 2, 2},
		want:        6,
		err:         "",
	},
}

var blocks = []struct {
	description string
	sum         int
	A           []int
	want        int
	err         string
}{
	{
		description: "Example",
		sum:         6,
		A:           []int{2, 1, 5, 1, 2, 2, 2},
		want:        3,
		err:         "",
	},
}
