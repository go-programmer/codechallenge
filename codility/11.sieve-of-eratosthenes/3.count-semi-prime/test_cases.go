package solution

// func Solution(N int, P []int, Q []int) []int
var semiprime = []struct {
	description string
	N           int
	P           []int
	Q           []int
	want        []int
	err         string
}{
	{
		description: "40",
		N:           40,
		P:           []int{24, 26, 6, 4, 1},
		Q:           []int{30, 32, 40, 40, 40},
		want:        []int{2, 1, 14, 15, 15},
		err:         "",
	},
	{
		description: "60",
		N:           60,
		P:           []int{58, 57, 27, 1, 4, 1, 1, 4, 16, 13, 17, 21, 8, 10},
		Q:           []int{60, 57, 27, 27, 4, 4, 26, 10, 20, 20, 17, 21, 31, 30},
		want:        []int{1, 1, 0, 10, 1, 1, 10, 4, 0, 2, 0, 1, 8, 7},
		err:         "",
	},
	{
		description: "4",
		N:           4,
		P:           []int{2, 4, 1, 3, 1, 2, 2},
		Q:           []int{4, 4, 4, 4, 1, 4, 2},
		want:        []int{1, 1, 1, 1, 0, 1, 0},
		err:         "",
	},
	{
		description: "7",
		N:           7,
		P:           []int{4, 1},
		Q:           []int{6, 7},
		want:        []int{2, 2},
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
