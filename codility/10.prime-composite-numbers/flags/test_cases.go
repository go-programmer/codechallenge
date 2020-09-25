package solution

var peaks = []struct {
	description       string
	A                 []int
	flag              int
	expectedPeaks     []int
	expectedDistances map[int]int
	err               string
}{
	{
		description:       "A has 4 peaks",
		A:                 []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
		flag:              3,
		expectedPeaks:     []int{1, 3, 5, 10},
		expectedDistances: map[int]int{2: 2, 5: 1},
		err:               "",
	},
	{
		description:       "There is no peak",
		A:                 []int{1, 2, 3, 4},
		expectedPeaks:     []int{},
		expectedDistances: map[int]int{},
		err:               "",
	},
	{
		description:       "There is one peak",
		A:                 []int{1, 3, 2},
		flag:              1,
		expectedPeaks:     []int{1},
		expectedDistances: map[int]int{},
		err:               "",
	},
	{
		description:       "Two peaks",
		A:                 []int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1},
		flag:              2,
		expectedPeaks:     []int{5, 7},
		expectedDistances: map[int]int{2: 1},
		err:               "",
	},
	{
		description:       "No peak",
		A:                 []int{3, 2, 1},
		flag:              0,
		expectedPeaks:     []int{},
		expectedDistances: map[int]int{},
		err:               "",
	},
	{
		description:       "One peak at the end",
		A:                 []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 10},
		flag:              1,
		expectedPeaks:     []int{10},
		expectedDistances: map[int]int{},
		err:               "",
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
	{
		description: "Four mountains, 3 flags",
		peaks:       []int{1, 3, 7, 9},
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
