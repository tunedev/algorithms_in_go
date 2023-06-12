package main

import (
	"fmt"

	linkedlist "github.com/tunedev/algorithms_in_go/linked_list"
)

func main() {
	l1 := linkedlist.NewLinkedList[int]()
	l1.AddFirst(12)
	l1.AddFirst(21)
	l1.AddFirst(23)
	l1.AddLast(2920)
	l1.AddLast(299)
	fmt.Println(l1)
	fmt.Println("size",l1.Size())
	fmt.Println(l1.ToArray())
l1.PrintMiddle()

}