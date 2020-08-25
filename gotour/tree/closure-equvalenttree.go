package main

import (
	"fmt"
	"os"

	"golang.org/x/tour/tree"
)

type ch chan int

func Walk(t *tree.Tree, ch chan int, done chan struct{}) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		select {
		case ch <- t.Value:
		case <-done:
			os.Goexit() // stop the current goroutine cleanly
		}
		walk(t.Right)
	}
	walk(t)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walker(t *tree.Tree, ch ch) {
	defer close(ch)

	var w func(t *tree.Tree)
	w = func(t *tree.Tree) {
		if t != nil {
			w(t.Right)
			ch <- t.Value
			w(t.Left)
		}
	}

	w(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(ch), make(ch)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for v := range c1 {
		if v != <-c2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(ch)

	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v, " ")
	}

	fmt.Println("Same?", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same?", Same(tree.New(1), tree.New(2)))
}
