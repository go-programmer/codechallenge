package solution

var choclates = []struct {
	description string
	N           int
	M           int
	want        int
	err         string
}{
	{
		description: "N10, M4",
		N:           10,
		M:           4,
		want:        5,
		err:         "",
	},
}
