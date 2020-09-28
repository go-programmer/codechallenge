package solution

var divisor = []struct {
	description string
	n           int
	want        int
	err         string
}{
	{
		description: "Highest divisor of 8",
		n:           8,
		want:        2,
		err:         "",
	},
	{
		description: "Highest divisor of 7",
		n:           7,
		want:        1,
		err:         "",
	},
	{
		description: "Highest divisor of 100",
		n:           100,
		want:        5, //1 2 4 5 10 20 25 50 100
		err:         "",
	},
	{
		description: "Highest divisor of 24",
		n:           24,
		want:        4,
		err:         "",
	},
	{
		description: "Highest divisor of 30",
		n:           30,
		want:        5,
		err:         "",
	},
	{
		description: "Highest divisor of 1000",
		n:           1000,
		want:        25,
		err:         "",
	},
	{
		description: "Highest divisor of 11",
		n:           11,
		want:        1,
		err:         "",
	},
	{
		description: "Perimeter of 36",
		n:           36,
		want:        4,
		err:         "",
	},
	{
		description: "Perimeter of 100000000",
		n:           100000000,
		want:        8000,
		err:         "",
	},
}

var perimeter = []struct {
	description string
	n           int
	want        int
	err         string
}{
	{
		description: "Perimeter of 30",
		n:           30,
		want:        22,
		err:         "",
	},

	{
		description: "Perimeter of 1",
		n:           1,
		want:        2,
		err:         "",
	},

	{
		description: "Perimeter of 36",
		n:           36,
		want:        24,
		err:         "",
	},
	{
		description: "Perimeter of 100000000",
		n:           100000000,
		want:        40000,
		err:         "",
	},
}

var square = []struct {
	description string
	n           int
	want        bool
	err         string
}{
	{
		description: "30 Not square",
		n:           30,
		want:        false,
		err:         "",
	},

	{
		description: "100000000 is square",
		n:           100000000,
		want:        true,
		err:         "",
	},
	{
		description: "36 is not square",
		n:           36,
		want:        false,
		err:         "",
	},
}
