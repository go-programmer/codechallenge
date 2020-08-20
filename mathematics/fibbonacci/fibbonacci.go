package fibbonacci

import (
	"fmt"
)

func fibbonacci(n int) []int {
	var slice []int
	x, y := 1, 1

	for i := 0; i < n; i++ {
		slice = append(slice, x)
		x, y = y, x+y
	}

	return slice
}

func fibonacciChan(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciSelect(out chan int, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case out <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

// Client clinet
func Client() {
	fibbonacciNumbers := 5
	c := make(chan int)
	switchChan := make(chan int)
	quit := make(chan int)

	// simple sequential fibbonacci
	fmt.Println("Serial")
	fibNum := fibbonacci(fibbonacciNumbers)
	for _, num := range fibNum {
		fmt.Println(num)
	}

	// With simple channel
	fmt.Println("With chan")
	go fibonacciChan(fibbonacciNumbers, c)
	for i := range c {
		fmt.Println(i)
	}

	// With Select
	fmt.Println("With switch chan")
	go func() {
		for i := 0; i < fibbonacciNumbers; i++ {
			fmt.Println(<-switchChan)
		}
		quit <- 0
	}()
	go fibonacciSelect(switchChan, quit)
}
