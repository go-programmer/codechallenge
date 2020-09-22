package nesting

var testCases = []struct {
	description string
	S           string
	want        int
	err         string
}{
	{
		description: "Example 1",
		S:           "(()(())())",
		want:        1,
		err:         "",
	},
	{
		description: "Example 2",
		S:           "([)()]",
		want:        1,
		err:         "",
	},
	{
		description: "Closed at middle",
		S:           "())(()",
		want:        0,
		err:         "",
	},
	{
		description: "No () brackets",
		S:           "{{{{",
		want:        1,
		err:         "",
	},
	{
		description: "One open-close pair",
		S:           "()",
		want:        1,
		err:         "",
	},
	{
		description: "First close",
		S:           ")(",
		want:        0,
		err:         "",
	},
	{
		description: "One unclosed",
		S:           "(()",
		want:        0,
		err:         "",
	},
}
