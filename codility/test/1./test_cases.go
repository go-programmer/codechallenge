package lesson

var sieve = []struct {
	description string
	S           string
	want       	string
	err         string
}{
	{
		description: "Case 1",
		S:           "00-44  48 5555 8361",
		want:        "004-448-555-583-61",
		err:         "",
	},
	{
		description: "Case 3",
		S:           "555372654",
		want:        "555-372-654",
		err:         "",
	},
	{
		description: "Case 2",
		S:           "0 - 22 1985--324",
		want:        "022-198-53-24",
		err:         "",
	},
	
	
}
