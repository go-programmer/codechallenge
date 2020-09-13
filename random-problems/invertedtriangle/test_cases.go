package invertedtriangle

var testCases = []struct {
	description string
	n           int
	want        string
	err         string
}{
	{
		description: "11 stars triangle.",
		n:           11,
		want:        `*********** *********   *******     *****       ***         *     `,
		err:         "",
	},
	{
		description: "5 stars trangle.",
		n:           5,
		want:        `***** ***   *  `,
		err:         "",
	},
	{
		description: "Should return empty string.",
		n:           2,
		want:        ``,
		err:         "",
	},
	{
		description: "One star.",
		n:           1,
		want:        `*`,
		err:         "",
	},
}
