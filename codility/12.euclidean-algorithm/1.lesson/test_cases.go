package solution

var gcdData = []struct {
	description string
	a           int
	b           int
	want        int
	err         string
}{
	{
		description: "GCD of 2 and 6",
		a:           2,
		b:           6,
		want:        2,
		err:         "",
	},
	{
		description: "GCD of 2 and 4",
		a:           2,
		b:           4,
		want:        2,
		err:         "",
	},
	{
		description: "GCD of 1 and 5",
		a:           5,
		b:           1,
		want:        1,
		err:         "",
	},
	{
		description: "GCD of 10 and 4",
		a:           10,
		b:           4,
		want:        2,
		err:         "",
	},
	{
		description: "GCD of 15 and 75",
		a:           15,
		b:           75,
		want:        15,
		err:         "",
	},
	{
		description: "GCD of 7 and 3",
		a:           7,
		b:           3,
		want:        1,
		err:         "",
	},
}
