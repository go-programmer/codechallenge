
// min 1
// max 483647
// Find longest sequence of zeros in binary representation of an integer.
// https://app.codility.com/programmers/lessons/1-iterations
func longestBinaryGap(N int) int {

	if N < 5 {
		return 0
	}

	// x := int(N)
	s := fmt.Sprintf("%b", N)
	fmt.Println(s)
	// c := []int(s)

	m, c := 0, 0
	for _, sv := range s {
		// fmt.Println(sv)

		// max counter m, internal counter c = 0
		// start internal counter ic when sv 49
		// if sv == 1; if c>m m = c; c=0
		// if sv == 0; c++

		if sv == 49 {
			if c > m {
				m = c
			}
			c = 0
		} else {
			c++
		}
		fmt.Printf("D:m:c %d:%d:%d \n", sv, m, c)
	}

	return m
}

func longestBinaryGapClient() {
	const name, age = "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.")
	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
	fmt.Println(s)

	// 561892
	longestBinaryGap(561892)
}