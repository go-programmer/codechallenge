package main

// iterations "codechallenge/codility/1.Iterations"
// "codechallenge/codility/challenges/stockslaundering"
// "codechallenge/codility/iterations"

import "fmt"

type Animal struct {
	count int
}

func main() {
	mainA()
}

func mainA() map[string]*Animal {
	// m := map[string]*Animal{"cat": &Animal{2}, "dog": &Animal{3}, "mouse": &Animal{5}}
	m := map[string]*Animal{}
	m["dog"] = &Animal{2}
	fmt.Printf("%#v\n", m["dog"])

	m["dog"].count = 4

	fmt.Printf("%#v", m["dog"])

	return m

}

// func callBinaryGap() {
// 	iterations.LongestBinaryGapClient()
// }

// func callStockLaundering() {
// 	stockslaundering.Client()
// }
