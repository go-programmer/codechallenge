package main

import (
	"container/list"
	"fmt"
)

func main() {
	mylist := list.New()
	mylist.PushFront(1)
	mylist.PushFront(2)
	mylist.PushBack(3)

	e := mylist.Front()
	for e != nil {
		fmt.Println(e.Value)

		e = e.Next()
	}
}
