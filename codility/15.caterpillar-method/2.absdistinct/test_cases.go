package solution

var numbers = []struct {
	description string
	A           []int
	want        int
	err         string
}{
	{
		description: "Provided example",
		A:           []int{-5, -3, -1, 0, 3, 6},
		want:        5,
		err:         "",
	},
}
