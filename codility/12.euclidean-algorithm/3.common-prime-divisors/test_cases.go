package solution

var countPrimeFactors = []struct {
	description string
	A           []int
	B           []int
	want        int
	err         string
}{
	{
		description: "15, 75 have a same prime divisors.",
		A:           []int{15, 10, 3, 4},
		B:           []int{75, 30, 5, 20},
		want:        1,
		err:         "",
	},
}

var primeDivisors = []struct {
	description string
	N           int
	M           int
	want        bool
	err         string
}{
	{
		description: "15, 75 have a same prime divisors.",
		N:           15,
		M:           75,
		want:        true,
		err:         "",
	},
	{
		description: "10, 30 does not have same prime divisors.",
		N:           10,
		M:           30,
		want:        false,
		err:         "",
	},
	{
		description: "20, 4 does not have same prime divisors.",
		N:           4,
		M:           20,
		want:        false,
		err:         "",
	},
}

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
