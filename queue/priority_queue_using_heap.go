package queue

// wrapper arround a heap
// it improves the Insert operation time complexity of the other implementation from O(n) to O(log n)
// But it worsens the time complexity of the Remove method from O(1) to O(log n)

import (
	h "github.com/tunedev/algorithms_in_go/heap"
)

type PriorityQueueWithHeap struct {
	heap *h.Heap
}

func NewPriorityQueuWithHeap() *PriorityQueueWithHeap {
	return &PriorityQueueWithHeap{
		heap: h.NewHeap(),
	}
}

func (p *PriorityQueueWithHeap) Enqueue(val int) {
	p.heap.Insert(val)
}

func (p *PriorityQueueWithHeap) Dequeue(val int) int {
	return p.heap.Remove()
}

func (p *PriorityQueueWithHeap) IsEmpty() bool {
	return p.heap.IsEmpty()
}