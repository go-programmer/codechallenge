// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// This is the provided solution at tour/solution
// Other sources of solutions:
// https://stackoverflow.com/questions/12224042/go-tour-exercise-equivalent-binary-trees
// https://blog.labix.org/2011/10/09/death-of-goroutines-under-control
// https://gist.github.com/inancgumus/d25d045b4cec43dcbb111e04980d396b#file-exercise_8_equivalent_binary_trees_with_goroutines-go
// https://stackoverflow.com/questions/12224042/go-tour-exercise-equivalent-binary-trees
// https://groups.google.com/g/golang-nuts/c/gva8SmZGmek
// https://gist.github.com/claw88/e8f3d913622fcdf4984cdd5ce7dcd3c5

package main

import (
	"fmt"

	"codechallenge/gotour/tree"
)

func walkImpl(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	// Need to close the channel here
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// NOTE: The implementation leaks goroutines when trees are different.
// See binarytrees_quit.go for a better solution.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	firstTree := tree.New(1)
	fmt.Println(firstTree)
	// if Same(tree.New(1), tree.New(1)) {
	// 	fmt.Println("PASSED")
	// } else {
	// 	fmt.Println("FAILED")
	// }

	// fmt.Print("tree.New(1) != tree.New(2): ")
	// if !Same(tree.New(1), tree.New(2)) {
	// 	fmt.Println("PASSED")
	// } else {
	// 	fmt.Println("FAILED")
	// }
}
