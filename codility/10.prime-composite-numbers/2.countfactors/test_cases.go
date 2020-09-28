package countfactors

var factors = []struct {
	description string
	n           int
	want        int
	err         string
}{
	{
		description: "24 has 8 factors",
		n:           24,
		want:        8,
		err:         "",
	},
	{
		description: "7 has 2 factors",
		n:           7,
		want:        2,
		err:         "",
	},
}
