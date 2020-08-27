package stockslaundering

import (
	"math/rand"
	"time"
)

// Client calls frogjump()
func Client() {

	var testcase string
	var want, got int

	var ml int       // ml=washing machine limit
	var ds, cs []int // ds=dirty socks, cl= clean socks

	// testcase = "1. Provided data"
	// ml = 2
	// ds = []int{1, 2, 1, 1}    // pairs=1	single=1,2
	// cs = []int{1, 4, 3, 2, 4} // pairs=4	single=1,2,3
	// want = 3
	// got = stocksLaundering(ml, ds, cs)
	// PrintFail(got, want, testcase)

	// testcase = "2. Pair does not exist."
	// ml = 1
	// ds = []int{1}
	// cs = []int{2}
	// want = 0
	// got = stocksLaundering(ml, ds, cs)
	// PrintFail(got, want, testcase)

	// testcase = "3. Two clean socks pairs and zero machine capacity."
	// ml = 0
	// cs = []int{2, 2, 1, 2, 3, 4, 5, 6, 7, 2, 8}
	// ds = []int{1}
	// want = 2
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "4. Machines capacity 0. No clean socks pair."
	// ml = 0
	// cs = []int{1, 2}
	// ds = []int{3}
	// want = 0
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "5. Single pair of clean socks."
	// ml = 1
	// cs = []int{3, 3}
	// ds = []int{4, 5, 6, 7, 8}
	// want = 1
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "6. When arent any pair of clean socks."
	// ml = 1
	// cs = []int{3}
	// ds = []int{5, 6, 7, 8}
	// want = 0
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "7. When there is no single clean sock."
	// ml = 1
	// cs = []int{3, 3}
	// ds = []int{5}
	// want = 1
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "8. When there is a dirty only one dirty sock."
	// ml = 5
	// cs = []int{1, 1, 3, 3, 7}
	// ds = []int{5}
	// want = 2
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "9. When machine limit is 1 but a pair of dirty socks exists."
	// ml = 1
	// cs = []int{3, 3}
	// ds = []int{5, 2, 2}
	// want = 1
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "10. When there is a one pair of clean and diirty socks."
	// ml = 9
	// cs = []int{3, 3, 1, 50}
	// ds = []int{1, 50}
	// want = 3
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "11. There are only dirty socks pair."
	// ml = 18
	// cs = []int{1, 3, 4}
	// ds = []int{2, 2, 50, 50, 19, 19, 45, 2, 2, 50, 50, 19, 19, 46, 7, 7}
	// want = 7
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "12. All socks are dirty."
	// ml = 14
	// cs = []int{1}
	// ds = []int{2, 2, 50, 50, 19, 19, 45, 45, 2, 2, 50, 50, 19, 19, 45, 7}
	// want = 7
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "13. Machine limit should not hinder already clean pair socks."
	// ml = 3
	// cs = []int{50, 2, 40}
	// ds = []int{40, 40, 40, 22, 22}
	// want = 2
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "14. All socks are dirty and machine limits decreases the clean paired socks."
	// ml = 6
	// cs = []int{1}
	// // ds = []int{50, 2, 2, 50, 50, 19, 19, 45, 2, 2, 50, 50, 19, 19, 45, 7}
	// ds = []int{2, 2, 50, 50, 19, 19}
	// want = 3
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "14. One dirty socks pair."
	// ml = 50
	// cs = []int{1, 4, 6, 9}
	// ds = []int{50, 2, 19, 45, 2, 7}
	// want = 1
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "15. Clean, dirty and mix pair"
	// ml = 5
	// cs = []int{1, 1, 4, 6, 9, 1, 1, 1}
	// ds = []int{50, 2, 19, 45, 2, 7, 9, 1}
	// want = 5
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "16. One dirty pair"
	// ml = 2
	// cs = []int{4}
	// ds = []int{50, 2, 19, 45, 2, 7, 9, 1}
	// want = 1
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "17. Two dirty pair"
	// ml = 23
	// cs = []int{4}
	// ds = []int{50, 2, 19, 45, 2, 7, 9, 1, 50}
	// want = 2
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "18. Random few colors"
	// ml = 23
	// cs = []int{1}
	// ds = []int{50, 13, 19}
	// want = 0
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	// testcase = "18. Random few colors"
	// ml = 50
	// rand.Seed(time.Now().UnixNano())
	// cs = rand.Perm(50)
	// ds = []int{2, 1, 1, 1}
	// want = 3
	// got = stocksLaundering(ml, cs, ds)
	// PrintFail(got, want, testcase)

	testcase = "19. Random max permutation with few repetitions"
	ml = 7
	rand.Seed(time.Now().UnixNano())
	cs = []int{1}
	ds = rand.Perm(15)
	ds = append(ds, 6, 8)
	// fmt.Println(ds)
	want = 3
	got = stocksLaundering(ml, cs, ds)
	PrintFail(got, want, testcase)
}
