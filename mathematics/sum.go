package mathematics

import (
	"fmt"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

// SumClient clent
func SumClient() {
	// runtime.GOMAXPROCS(4)
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
