package reversestring

var testCases = []struct {
	description string
	word        string
	want        string
	err         string
}{

	{
		description: "Pallendrome `ciric`",
		word:        `ciric`,
		want:        `cirica`,
		err:         "",
	},
	{
		description: "Reverse `anit`",
		word:        `anit`,
		want:        `tinaa`,
		err:         "",
	},
}
