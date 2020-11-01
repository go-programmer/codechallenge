package solution

var numbers = []struct {
	description string
	A           []int
	M        	int
	want        int
	err         string
}{
	{
		description: "Provided example",
		A:          []int{3,4,5,5,2},
		M:			6,
		want:       9,
		err:        "",
	},
}
