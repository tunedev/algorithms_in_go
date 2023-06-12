package queue

import (
	"errors"
	"fmt"
)

// Queue is a data structure that allows you process jobs based on the order we received them
// Common operations with their runtime:
// 	Enqueue: O(1)
// 	Dequeue: O(1)
// 	Peek: O(1)
// 	IsEmpty: O(1)
// 	IsFull: O(1)

type ArrayQueue struct{
	items []int
	rear int
	front int
	count int
	size int
}

var ErrMaxCapacity = errors.New("queue is at maximum capacity")
var ErrInvalidOp = errors.New("invalid state, queue is empty")

func NewQueue(size int) *ArrayQueue {
	return &ArrayQueue{
		items: make([]int, size),
		size: size,
	}
}

func (q *ArrayQueue) Enqueue(val int) {
	if q.IsFull() {
		panic(ErrMaxCapacity)
	}
	rear := q.rear % q.size
	q.items[rear] = val
	q.rear += 1
	q.count += 1
}

func (q *ArrayQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("invalid state, queue is empty")
	}
	derivedFront := q.front % q.size
	result := q.items[derivedFront]
	q.items[derivedFront] = 0
	q.front += 1;
	q.count -= 1;
	return result;
}

func (q *ArrayQueue) Peek() int {
	if q.IsEmpty() {
		panic(ErrInvalidOp)
	}
	return q.items[q.front % q.size];
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.count == 0
}

func (q *ArrayQueue) IsFull() bool {
	return q.count == q.size
}

func (q *ArrayQueue) String() string {
	result := []int{}
	for i := 0; i < q.count; i++ {
		result = append(result, q.items[(q.front + i) % q.size])
	}
	return fmt.Sprintln(result)
}

