package frogjump

import "fmt"

// Problem:
// Count minimal number of jumps from position X to Y.
// Write a function:
// func Solution(X int, Y int, D int) int
// that, given three integers X, Y and D,
// returns the minimal number of jumps D from
// position X to a position equal to or greater
// than Y.
// https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp

// Solution scores 100% in Codility Arrays CyclicRotation
// Detected time complexity: O(N)

func frogJump(X int, Y int, D int) int {

	if X == Y {
		return 0
	}

	if Y == X+D {
		return 1
	}

	mod := (Y - X) % D
	remainder := (Y - X) / D

	if mod > 0 {
		remainder++
	}

	return remainder
}

// Client calls frogjump()
func Client() {
	var want, got int

	// case 1
	want = 3
	got = frogJump(10, 85, 30)
	// 10 + 30 * steps >= 85
	// steps = ( 85 - 10 ) / 30 = 75/30 = 2. Here remainder is more than 1, so increase
	fmt.Printf("1 g:%d	w:%d\n", got, want)
	if got == want {
		fmt.Println("1 Ok")
	}

	// case 2
	want = 8
	got = frogJump(5, 85, 10) // 5 + 10 * 8
	fmt.Printf("2 g:%d	w:%d\n", got, want)
	// 5 + 10 * steps >= 85
	// steps = ( 85 - 5 ) / 10 = 80/10 = 8
	if got == want {
		fmt.Println("2 Ok")
	}

	// case 3 X==Y
	want = 0
	got = frogJump(85, 85, 30)
	fmt.Printf("3 g:%d	w:%d\n", got, want)
	if got == want {
		fmt.Println("3 Ok")
	}

	// case 4
	// Starting point is 0
	want = 10
	got = frogJump(0, 10, 1)
	fmt.Printf("4 g:%d	w:%d\n", got, want)
	if got == want {
		fmt.Println("4 Ok")
	}

	// case 5
	// Jump == Step
	want = 1
	got = frogJump(1, 10, 10)
	fmt.Printf("5 g:%d	w:%d\n", got, want)
	if got == want {
		fmt.Println("5 Ok")
	}

}
