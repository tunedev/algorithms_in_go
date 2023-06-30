package heap

import "fmt"

// A heap is a special type of tree with two properties
// 	1) Complete: Every level is filled from left to right(except potentially the last level)
//  2) Heap property: MaxHeap or MinHeap:
// 			MaxHeap: val of every node is greater or equal to its children
// 			MinHeap: val of every node is less or equal to its children

// Heap applications:
// - Sorting(HeapSort)
// - Graph algorithms (shortest path)
// - Priority queues
// - Finding the Kth smallest/largest value

/* in heap we can only delete the root node,
then we promote the right most node in the leaf nodes to the root,
then bubble down from there, by bubbling up the node with the largest
value of its imediate children
*/

// implementing a heap using array to reduce the memory foot print
// left = parent * 2 + 1
// right = parent * 2 + 2
// parent = (index -1) / 2

type Heap struct {
	items []int
	size int
}

func NewHeap() *Heap {
	return &Heap{
		items: []int{},
	}
}

func (h *Heap) Insert(val int) {
	h.items = append(h.items, val)
	h.size ++
	h.bubleUp()
}

func (h *Heap) bubleUp() {
	index := h.size - 1
	for (index > 0 && h.items[index] > h.getParent(index)) {
		h.swap(index, h.getParentIndex(index))
		index = h.getParentIndex(index)
	}
}

func (h *Heap) swap(first, second int) {
	temp := h.items[first]
	h.items[first] = h.items[second]
	h.items[second] = temp
}

func (h *Heap) getParent(index int) int {
	return h.items[h.getParentIndex(index)]
}
func (h *Heap) getParentIndex(index int) int {
	return (index -1) / 2
}

func (h *Heap) Remove() int {
	if h.IsEmpty(){
		panic("Heap is Empty, invalid state exception")
	}
	h.size--
	res := h.items[0]
	h.items[0] = h.items[h.size]

	h.bubbleDown()

	h.items = h.items[0:h.size]
	return res
}

func (h *Heap) bubbleDown() {
	index := 0
	for index <= h.size && !h.isValidParent(index) {
		highestChildIndex := h.getHighestChildsIndex(index)
		h.swap(index, highestChildIndex)
		index = highestChildIndex
	}
}

func (h *Heap) IsEmpty() bool {
	return  h.size == 0
}

func (h *Heap) isValidParent(parentIndex int) bool {
	if !h.hasLeftChild(parentIndex) {
		return true
	}

	isValid := h.items[parentIndex] >= h.getLeftChild(parentIndex)

	if h.hasRightChild(parentIndex) {
		isValid = isValid && h.items[parentIndex] >= h.getRightChild(parentIndex)
	}

	return isValid
}

func (h *Heap) getHighestChildsIndex(index int) int {
	if !h.hasLeftChild(index) {
		return index
	}

	if !h.hasRightChild(index) {
		return h.leftChildIndex(index)
	}

	if h.getLeftChild(index) >= h.getRightChild(index) {
			return h.leftChildIndex(index)
		} 
			return h.rightChildIndex(index)
		
}

func (h *Heap) hasLeftChild(index int) bool {
	return h.leftChildIndex(index) < h.size
}

func (h *Heap) hasRightChild(index int) bool {
	return h.rightChildIndex(index) < h.size
}

func (h *Heap) getLeftChild(parentIndex int) int {
	return h.items[h.leftChildIndex(parentIndex)]
}

func (h *Heap) getRightChild(parentIndex int) int {
	return h.items[h.rightChildIndex(parentIndex)]
}

func (h *Heap) leftChildIndex(parentIndex int) int {
	return parentIndex * 2 + 1
}

func (h *Heap) rightChildIndex(parentIndex int) int {
	return parentIndex * 2 + 2
}

func (h *Heap) String() string {
	var result string

	for i, v := range h.items {
		result += fmt.Sprintf("%v = %v\n", i, v)
	}
	return result
}

