package lesson

var sieve = []struct {
	description string
	n           int
	want        []int
	err         string
}{
	{
		description: "12 is prime",
		n:           12,
		want:        []int{1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1},
		err:         "",
	},
}

var smallestPrimeFactor = []struct {
	description string
	n           int
	want        []int
	err         string
}{
	{
		description: "Smallest prime factors of twelve numbers.",
		n:           12,
		want:        []int{0, 0, 0, 0, 2, 0, 2, 0, 2, 3, 2, 0, 2},
		err:         "",
	},
	{
		description: "Smallest prime factors of 25 numbers.",
		n:           25,
		want:        []int{0, 0, 0, 0, 2, 0, 2, 0, 2, 3, 2, 0, 2, 0, 2, 3, 2, 0, 2, 0, 2, 3, 2, 0, 2, 5},
		err:         "",
	},
	{
		description: "Smallest prime factors of 26 numbers.",
		n:           26,
		want:        []int{0, 0, 0, 0, 2, 0, 2, 0, 2, 3, 2, 0, 2, 0, 2, 3, 2, 0, 2, 0, 2, 3, 2, 0, 2, 5, 2},
		err:         "",
	},
}

var primeFactor = []struct {
	description string
	n           int
	want        []int
	err         string
}{
	{
		description: "Prime factors of twelve numbers.",
		n:           12,
		want:        []int{2, 2, 3},
		err:         "",
	},
	{
		description: "Prime factors of 25.",
		n:           25,
		want:        []int{5, 5},
		err:         "",
	},
	{
		description: "Prime factors of 100.",
		n:           100,
		want:        []int{2, 2, 5, 5},
		err:         "",
	},

	{
		description: "Prime factors of 100000.",
		n:           100000,
		want:        []int{2, 2, 2, 2, 2, 5, 5, 5, 5, 5},
		err:         "",
	},
	{
		description: "Prime factors of 26.",
		n:           26,
		want:        []int{2, 13},
		err:         "",
	},
	{
		description: "Prime factors of 16.",
		n:           16,
		want:        []int{2, 2, 2, 2},
		err:         "",
	},
}
