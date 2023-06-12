package queue

import "fmt"

type Node[T any] struct{
	weight int
	data T
}

type PriorityQueue[T any] struct {
	items []Node[T]
	count int
	size int
}

func NewPriorityQueue[T any](size int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: make([]Node[T], size),
		size: size,
	}
}

func (p *PriorityQueue[T]) Add(item T, priority int) error {
	if p.IsFull() {
		return ErrMaxCapacity
	}
	newNode := Node[T]{
		data: item,
		weight: priority,
	}
	var i int
	for i = p.count - 1; i >= 0; i -= 1 {
		if p.items[i].weight >= priority {
			p.items[i+1] = p.items[i]
		} else{
			break
		}
	}
	fmt.Println("adding in index ====>>", i + 1)
	p.items[i + 1] = newNode
	p.count += 1
	return nil
}

func (p *PriorityQueue[T]) Remove() (T, error) {
	if p.IsEmpty() {
		return (&Node[T]{}).data, ErrInvalidOp
	}
	p.count -= 1
	return p.items[p.count].data, nil
}

func (p *PriorityQueue[T]) IsEmpty() bool {
	return p.count == 0
}

func (p *PriorityQueue[T]) IsFull() bool {
	return p.count == p.size
}

func (p *PriorityQueue[T]) String() string {
	result := make([]T, p.count)
	for i := p.count - 1; i >= 0; i -= 1 {
		result[p.count - i - 1] = p.items[i].data
	}
	return fmt.Sprintln(result)
}