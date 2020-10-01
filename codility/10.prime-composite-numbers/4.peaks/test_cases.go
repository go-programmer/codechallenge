package solution

var blocks = []struct {
	description string
	A           []int
	blocks      int
	err         string
}{
	{
		description: "A has one element",
		A:           []int{3},
		blocks:      0,
		err:         "",
	},
	{
		description: "A has two element",
		A:           []int{1, 3},
		blocks:      0,
		err:         "",
	},
	{
		description: "Odd length with a peak, returns 1 block",
		A:           []int{1, 3, 2},
		blocks:      1,
		err:         "",
	},
	{
		description: "Len 3, Picks 0, Blocks 0",
		A:           []int{1, 2, 2},
		blocks:      0,
		err:         "",
	},
	{
		description: "Odd length with a peak, returns 1 block",
		A:           []int{1, 3, 2, 5, 5},
		blocks:      1,
		err:         "",
	},
	{
		description: "A has 3 blocks",
		A:           []int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
		blocks:      3,
		err:         "",
	},
	{
		description: "2 peaks but 2 block",
		A:           []int{1, 2, 1, 2, 1, 2},
		blocks:      2,
		err:         "",
	},
	{
		description: "2 peaks but 1 block",
		A:           []int{1, 2, 1, 1, 1},
		blocks:      1,
		err:         "",
	},
	{
		description: "No peaks, No block",
		A:           []int{1, 2, 3, 4, 5},
		blocks:      0,
		err:         "",
	},
	{
		description: "3 peaks, 3 blocks",
		A:           []int{0, 1, 0, 0, 1, 0, 0, 1, 0},
		blocks:      3,
		err:         "",
	},
	{
		description: "3 peaks, 3 blocks",
		A:           []int{0, 1, 0, 0, 1, 0, 0, 1, 0},
		blocks:      3,
		err:         "",
	},

	{
		description: "4 peaks, 3 blocks",
		// 	     0  1  2  3  4  5  6  7  8  9  10 11
		A: []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
		// 		 	1     2     3              4
		// 12=3*4; 3 blocks with 4 elements, or 4 blocks with 3 elements
		// 4 Blocks of size 3, is not possible.
		// 3 blocks of size 4 is possible.
		blocks: 3,
		err:    "",
	},
	{
		description: "4 peaks, 3 blocks",
		A:           []int{1, 5, 3, 4, 3, 4, 1, 4, 3, 4, 6, 2},
		blocks:      4,
		err:         "",
	},
	{
		description: "4 peaks, 4 blocks",
		A:           []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2, 3, 1, 1, 1},
		blocks:      4,
		err:         "",
	},
	{
		description: "3 peaks, 4 blocks",
		A:           []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4},
		// Len A 10 Picks 3, factors are 2*5 and 1*10
		// For 1*10, one block only.
		// Either 5 blocks or 2 Len but picks is just 3, so not possible
		// Or 2 blocks 5 length, So 2 blocks as picks len is > 3
		blocks: 2,
		err:    "",
	},
	{
		description: "4 peaks 2 blocks, ",
		A:           []int{1, 5, 1, 5, 1, 5, 1, 5, 1, 5},
		// Len A 10 Picks 3, factors are 2*5 and 1*10
		// For 1*10, one block only.
		// Either 5 blocks or 2 Len but picks is just 3, so not possible
		// Or 2 blocks 5 length, So 2 blocks as picks len is > 3
		blocks: 2,
		err:    "",
	},
	{
		description: "Lenght 10, Peaks 4, Blocks 2 ",
		A:           []int{5, 1, 5, 1, 5, 1, 5, 1, 5, 1},
		// Len A 10 Picks 3, factors are 2*5 and 1*10
		// For 1*10, one block only.
		// Either 5 blocks or 2 Len but picks is just 3, so not possible
		// Or 2 blocks 5 length, So 2 blocks as picks len is > 3
		blocks: 2,
		err:    "",
	},
	{
		description: "Lenght 10, Peaks 4, Blocks 2 ",
		A:           []int{1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5},
		// Len A 10 Picks 3, factors are 2*5 and 1*10
		// For 1*10, one block only.
		// Either 5 blocks or 2 Len but picks is just 3, so not possible
		// Or 2 blocks 5 length, So 2 blocks as picks len is > 3
		blocks: 4,
		err:    "",
	},
}

var peaks = []struct {
	description string
	A           []int
	peaks       []int
	err         string
}{
	{
		description: "A has 4 peaks",
		A:           []int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
		peaks:       []int{3, 5, 10},
		err:         "",
	},
	{
		description: "A has 4 peaks",
		A:           []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
		peaks:       []int{1, 3, 5, 10},
		err:         "",
	},
	{
		description: "There is no peak",
		A:           []int{1, 2, 3, 4},
		peaks:       []int{},
		err:         "",
	},
	{
		description: "There is one peak",
		A:           []int{1, 3, 2},
		peaks:       []int{1},
		err:         "",
	},
	{
		description: "Two peaks",
		A:           []int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1},
		peaks:       []int{5, 7},
		err:         "",
	},
	{
		description: "No peak",
		A:           []int{3, 2, 1},
		peaks:       []int{},
		err:         "",
	},
	{
		description: "One peak at the end",
		A:           []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 10},
		peaks:       []int{10},
		err:         "",
	},
	{
		description: "One peak at odd len",
		A:           []int{2, 6, 4, 5, 6, 5, 1},
		peaks:       []int{1},
		err:         "",
	},
	{
		description: "Two peaks",
		A:           []int{1, 2, 1, 2, 1, 2},
		peaks:       []int{1, 3},
		err:         "",
	},
	{
		description: "Two peaks",
		A:           []int{0, 1, 0, 0, 1, 0, 0, 1, 0},
		peaks:       []int{1, 4, 7},
		err:         "",
	},
}

var factors = []struct {
	description string
	A           int
	factorials  []int
	err         string
}{
	{
		description: "50000",
		A:           50000,
		factorials:  []int{5},
		err:         "",
	},
}
