package brackets

var testCases = []struct {
	description string
	S           string
	want        int
	err         string
}{
	{
		description: "Example 1",
		S:           "{[()()]}",
		want:        1,
		err:         "",
	},
	{
		description: "Example 2",
		S:           "([)()]",
		want:        0,
		err:         "",
	},
	{
		description: "Failed case",
		S:           "())(()",
		want:        0,
		err:         "",
	},
	{
		description: "Failed case",
		S:           "{{{{",
		want:        0,
		err:         "",
	},
	{
		description: "Failed case",
		S:           "()",
		want:        1,
		err:         "",
	},
	{
		description: "Failed case",
		S:           "+)(",
		want:        0,
		err:         "",
	},
	{
		description: "Failed case",
		S:           "+ )( + ()",
		want:        0,
		err:         "",
	},
}
