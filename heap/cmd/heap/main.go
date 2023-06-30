package main

import (
	"fmt"

	h "github.com/tunedev/algorithms_in_go/heap"
)

func main() {
	// heap := h.NewHeap()

	// heap.Insert(10)
	// heap.Insert(5)
	// heap.Insert(17)
	// heap.Insert(4)
	// heap.Insert(22)
	// fmt.Println(heap.Remove())
	// heap.Insert(60)
	// fmt.Println("Done")

	fmt.Println(heapify2([]int{5,3,8,4,1,2}))
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

// My implementation, reminicing on the actual struct
func heapify(items []int) []int {
	leftChildIndex := func(index int) int {
		return index * 2 + 1
	}
	rightChildIndex := func(index int) int {
		return index * 2 + 2
	}

	hasLeftChild := func(index int) bool {
		return leftChildIndex(index) < len(items)
	}

	hasRightChild := func(index int) bool {
		return rightChildIndex(index) < len(items)
	}

	getLeftChild := func(index int) int {
		return items[leftChildIndex(index)]
	}
	getRightChild := func(index int) int {
		return items[rightChildIndex(index)]
	}

	isValidParent := func(index int) bool {
		if !hasLeftChild(index) {
			return true
		}
		isValid := items[index] > getLeftChild(index)

		if hasRightChild(index) {
			isValid = isValid && items[index] > getRightChild(index)
		}
		return isValid
	}
	getMaxChildsIndex := func(index int) int {
		if getLeftChild(index) >= getRightChild(index) {
			return leftChildIndex(index)
		}
		return rightChildIndex(index)
	}

	bubbleDown := func(index int) {
		for index < len(items) && !isValidParent(index) {
			maxChildIndex := getMaxChildsIndex(index)
			items[index], items[maxChildIndex] = items[maxChildIndex], items[index]
			index = maxChildIndex
		}
	}

	for i := 0; i < len(items); i++ {
		bubbleDown(i)
	}
	return items
}

func heapify2(items []int) []int {
	// for i := 0; i < len(items); i++ {
	// 	heapifyHelper(&items, i)
	// }
	indexOfLastParent := len(items)/2 -1
	// // Suboptimal approach as it iterates over all the items to the last parent
	// for i := 0; i < indexOfLastParent; i++ {
	// 	heapifyHelper(&items, i)
	// }
	// // Optimal because it starts from the last parent, which on its way up would have been heapified properly
	for i := indexOfLastParent; i >= 0; i-- {
		heapifyHelper(&items, i)
	}
	return items
}

func heapifyHelper(items *[]int, index int) {
	largerIndex := index
	itemSize := len(*items)

	leftChildIndex := index * 2 + 1
	if leftChildIndex < itemSize && (*items)[leftChildIndex] > (*items)[index] {
		largerIndex = leftChildIndex
	}

	rightChildIndex := index * 2 + 2
	if rightChildIndex < itemSize && (*items)[rightChildIndex] > (*items)[leftChildIndex] {
		largerIndex = rightChildIndex
	}

	if largerIndex == index {
		return
	}

	(*items)[largerIndex], (*items)[index] = (*items)[index], (*items)[largerIndex]
	heapifyHelper(items, largerIndex)
}
