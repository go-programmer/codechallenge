package lesson

var primality = []struct {
	description string
	n           int
	want        bool
	err         string
}{
	{
		description: "7 is prime",
		n:           7,
		want:        true,
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
		description: "8 has 4 divisors",
		n:           8,
		want:        4,
		err:         "",
	},
	{
		description: "7 has 1 divisor",
		n:           7,
		want:        2,
		err:         "",
	},
}

var coins = []struct {
	description string
	n           int
	want        int
	err         string
}{
	{
		description: "1..9 coins has 1,4,9 flipped",
		n:           9,
		want:        3,
		err:         "",
	},
}
