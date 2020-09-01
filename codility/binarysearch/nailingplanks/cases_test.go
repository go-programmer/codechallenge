package nailingplanks

var nailedCases = []struct {
	description string
	low         int
	high        int
	nails       []int
	want        int
	err         string
}{
	{
		description: "Five nails nailed.",
		low:         1,
		high:        5,
		nails:       []int{1, 2, 3, 4, 5},
		want:        5,
		err:         "",
	},
	{
		description: "Two nails nailed.",
		low:         1,
		high:        5,
		nails:       []int{1, 2},
		want:        2,
		err:         "",
	},
	{
		description: "No nails nailed.",
		low:         1,
		high:        5,
		nails:       []int{6, 7, 8},
		want:        -1,
		err:         "",
	},
	{
		description: "One nail after > high.",
		low:         1,
		high:        5,
		nails:       []int{6, 7, 8, 5},
		want:        -1,
		err:         "",
	},
	{
		description: "When nail is less than low",
		low:         3,
		high:        5,
		nails:       []int{1, 2},
		want:        -1,
		err:         "",
	},
	{
		description: "One nails nailed",
		low:         2,
		high:        5,
		nails:       []int{1, 2},
		want:        1,
		err:         "",
	},
	{
		description: "When one nails is between range",
		low:         3,
		high:        5,
		nails:       []int{1, 2, 4},
		want:        1,
		err:         "",
	},
}

var testCases = []struct {
	description string
	plankA      []int
	plankB      []int
	nails       []int
	want        int
	err         string
}{
	{
		description: "Single plank with a valid nail",
		plankA:      []int{6},
		plankB:      []int{9},
		nails:       []int{6},
		want:        1,
		err:         "",
	},

	{
		description: "Four nails nailed!",
		plankA:      []int{1, 4, 5, 8},
		plankB:      []int{4, 5, 9, 10},
		nails:       []int{4, 6, 7, 10, 2},
		want:        4,
		err:         "",
	},
}
