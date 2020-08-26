package tennistournament

import (
	"fmt"
)

// Client calls frogjump()
func Client() {
	// P and C are integers within the range [1..30,000].
	var testcase string
	var want, got int

	testcase = "1. When there is less two players"
	want = 0
	got = totalGamesPlayed(1, 2)
	PrintFail(got, want, testcase)

	testcase = "2. When there is only one court available"
	want = 1
	got = totalGamesPlayed(3, 1)
	PrintFail(got, want, testcase)

	testcase = "3. When P is double C"
	want = 5
	got = totalGamesPlayed(10, 5)
	PrintFail(got, want, testcase)

	testcase = "4. When C is double P"
	want = 5
	got = totalGamesPlayed(5, 10)
	PrintFail(got, want, testcase)

	testcase = "5. When P is twice C"
	want = 6
	got = totalGamesPlayed(12, 6)
	PrintFail(got, want, testcase)

	testcase = "6. When P>C"
	want = 7
	got = totalGamesPlayed(14, 13)
	PrintFail(got, want, testcase)

	testcase = "7. When P=C"
	want = 7
	got = totalGamesPlayed(15, 15)
	PrintFail(got, want, testcase)

	testcase = "8. When C > P"
	want = 4
	got = totalGamesPlayed(9, 25)
	PrintFail(got, want, testcase)

	testcase = "8. When P > C"
	want = 3
	got = totalGamesPlayed(10, 3)
	PrintFail(got, want, testcase)

	testcase = "9. When there are only two players"
	want = 1
	got = totalGamesPlayed(2, 31)
	PrintFail(got, want, testcase)

}

func totalGamesPlayed(P int, C int) int {

	if P < 2 {
		return 0
	}

	if C == 1 || P == 2 {
		return 1
	}

	if P >= C {

		if P > C && P%C == 0 {
			return C
		}

		if P > C {
			if P/2 > C {
				return C
			}
		}

		return P / 2
	} else {

		if C%P == 0 {
			return P
		}

		return P / 2
	}
}

// PrintFail Print the failed case
func PrintFail(got int, want int, failMsg string) {
	if got != want {
		fmt.Printf("FAIL %v		g:%d w:%d\n", failMsg, got, want)
	}
}
