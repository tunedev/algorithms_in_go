package main

import (
	"fmt"

	bst "github.com/tunedev/algorithms_in_go/binary_tree"
)

func main() {
	t := bst.NewBST()

	t.Insert(7)
	t.Insert(4)
	t.Insert(1)
	t.Insert(6)
	t.Insert(9)
	t.Insert(8)
	t.Insert(10)
	fmt.Println(t)
	t.PrintTreeElementInOrder()
}