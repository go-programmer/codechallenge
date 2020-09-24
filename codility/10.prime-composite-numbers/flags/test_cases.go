package solution

var peaks = []struct {
	description string
	A           []int
	want        []int
	err         string
}{
	{
		description: "A has 4 peaks",
		A:           []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
		want:        []int{1, 3, 5, 10},
		err:         "",
	},
	{
		description: "There is no peak",
		A:           []int{1, 2, 3, 4},
		want:        []int{},
		err:         "",
	},
	{
		description: "There is one peak",
		A:           []int{1, 3, 2},
		want:        []int{1},
		err:         "",
	},
	{
		description: "Two peaks",
		A:           []int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1},
		want:        []int{5, 7},
		err:         "",
	},
	{
		description: "No peak",
		A:           []int{3, 2, 1},
		want:        []int{},
		err:         "",
	},
}

var flags = []struct {
	description string
	peaks       []int
	want        int
	err         string
}{
	{
		description: "4 peaks, 3 flags",
		peaks:       []int{1, 3, 5, 10},
		want:        3,
		err:         "",
	},
	{
		description: "2 peaks, 2 flags",
		peaks:       []int{5, 7},
		want:        2,
		err:         "",
	},
}

var peakandflag = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "A has 4 peaks, 3 flags",
		A:           []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
		want:        3,
		err:         "",
	},
	{
		description: "One peak, one flag",
		A:           []int{1, 3, 2},
		want:        1,
		err:         "",
	},
	{
		description: "No peak",
		A:           []int{3, 2, 1},
		want:        0,
		err:         "",
	},
}
