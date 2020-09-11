package discintersections

type point struct {
	x int
	y int
}

func solution(A []int) int {

	// graph := make(map[point]bool)
	// point0 := point{0, 1}
	// graph[point0] = true

	// graph := make(map[point]bool)

	// fmt.Println(graph[point0])

	c, l, h := 0, 0, 0
	m := len(A) - 1

	// For 0, low and high are not needed.
	if A[0] > m {
		c = m
	} else {
		c = A[0]
	}

	// fmt.Printf("(%v,%v):%v\n", 0, A[0], c)

	// for k, v := range A[1:] {
	for i := 1; i <= m; i++ {
		// fmt.Println(k, v)

		// set low
		if i < A[i] {
			l = 0
		} else {
			l = i - A[i]
		}

		// set high
		h = i + A[i]
		if h > m {
			h = m
		}

		for l <= h {

			// check if previous index covers current
			if l < i && A[l] > i {
				c++
			}

			if l > i {
				c++
			}

			l++
		}

		// fmt.Printf("(%v,%v):%v\n", i, A[i], c)

	}

	if c > 9999999 {
		return -1
	}

	return c
}
