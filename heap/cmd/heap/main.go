package main

import (
	"fmt"

	h "github.com/tunedev/algorithms_in_go/heap"
)

func main() {
	heap := h.NewHeap()

	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(17)
	heap.Insert(4)
	heap.Insert(22)
	fmt.Println(heap.Remove())
	heap.Insert(60)
	fmt.Println("Done")
}

func sortIntArrUsingHeap(items []int) ([]int, []int) {
	desc := make([]int, len(items))
	asc := make([]int, len(items))
	heap := h.NewHeap()
	for _, v := range items {
		heap.Insert(v)
	}
	// DESC
	for i := 0; i < len(desc); i++ {
		desc[i] = heap.Remove()
	}

	// ASC
	for i := len(asc) - 1; i >= 0; i-- {
		asc[i] = heap.Remove()
	}
	return desc, asc
}
