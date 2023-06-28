package main

import (
	"fmt"
	a "github.com/tunedev/algorithms_in_go/avl_trees"
)

func main() {
	t := a.NewAVLTree()

	t.Insert(10)
	t.Insert(8)
	t.Insert(6)
	fmt.Println(t)
}
