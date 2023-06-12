package linkedlist

import (
	"errors"
	"fmt"
)

// definition: Linked List are sequence of object that link to each other, what this implies is that each object holds a refernce to the next object
// 		The first node is called the head while the last node i called the tail
// Runtime complexities:
// 	Lookup:
// 		By Value: O(n)
// 		By index: O(n)
// 	Insert && Delete
// 		At the end: O(1) for Delete: O(n)
// 		At the begining: O(1)
// 		Somewhere in the middle or by index: O(n)

type Node[T comparable] struct {
	val T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	first *Node[T]
	last *Node[T]
	count int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) String() string {
	result := ""
	current := l.first
	indexCount := 1
	for current != nil {
		result += fmt.Sprintf("Element: #%v has value: %v\n", indexCount, current.val)
		current = current.next
		indexCount += 1
	}
	return result
}

func (l *LinkedList[T]) AddFirst(val T) {
	newNode := &Node[T]{
		val: val,
	}
	if l.IsEmpty() {
		l.first, l.last = newNode, newNode;
	}else {
		newNode.next = l.first
		l.first = newNode
	}
	l.count += 1
}
func (l *LinkedList[T]) AddLast(val T) {
	newNode := &Node[T]{
		val: val,
	}
	if l.IsEmpty() {
		l.first, l.last = newNode, newNode;
	}else {
		l.last.next = newNode
		l.last = newNode
	}
	l.count += 1
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.first == nil && l.last == nil
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}
	if l.first == l.last {
		l.first, l.last = nil, nil
	}else {
		nextFirst := l.first.next;
		l.first.next = nil
		l.first = nextFirst
	}
	l.count -= 1
}
func (l *LinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}
	if l.first == l.last {
		l.first, l.last = nil, nil
	}else {
		secondToLast := l.getPreviousNode(l.last)
		secondToLast.next = nil
		l.last = secondToLast
	}
	l.count -= 1
}

func (l *LinkedList[T]) getPreviousNode(node *Node[T]) *Node[T] {
	current := l.first;
	for current != nil {
		if current.next == node {
			return current
		}
		current = current.next
	}
	return nil
}

func (l *LinkedList[T]) Contains(val T) bool {
	if l.IsEmpty() {
		return false
	}
	return l.IndexOf(val) != -1
}

func (l *LinkedList[T]) IndexOf(val T) int {
	if l.IsEmpty() {
		return -1
	}
	indexCount := 0
	current := l.first;

	for current != nil {
		if current.val == val {
			return indexCount
		}
		current = current.next
		indexCount += 1
	}
	return -1
}

func (l *LinkedList[T]) Remove(val T) error {
	if l.IsEmpty() {
		return errors.New("invalid state for operation")
	}

	if l.first == l.last {
		l.first, l.last = nil, nil
		return nil
	}

	current := l.first
	for current != nil {
		if current.val == val {
			break;
		}
		current = current.next
	}

	if current == nil {
		return errors.New("key does not exist")
	}
	
	prev := l.getPreviousNode(current)
	fmt.Println("wetin dey sup sef ====>>>", prev, current)
	prev.next = current.next
	current.next = nil
	return nil
}

func (l *LinkedList[T]) Size() int {
	return l.count
}

func (l *LinkedList[T]) ToArray() []T {
	result := make([]T, l.count)
	current := l.first
	index := 0
	for current != nil {
		result[index] = current.val
		index += 1
		current = current.next
	}
	return result
}

func (l *LinkedList[T]) Reverse() {
	if l.IsEmpty() {
		return
	}
	prev, current := l.first, l.first.next
	for current != nil {
		temp := current.next;
		current.next = prev;
		prev = current;
		current = temp
	}

	l.first = prev
	l.last = l.first
	l.last = nil
}

func (l *LinkedList[T]) GetKthNodeFromEnd(k int) T {
	if l.IsEmpty() || k > l.count{
		panic("Illigal arg")
	}
	slow, fast := l.first, l.first;

	for i := 0; i < k - 1; i++ {
		fast = fast.next
	}
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	return slow.val
}

func (l *LinkedList[T]) PrintMiddle() {
	if l.IsEmpty() {
		return
	}
	fast, slow := l.first, l.first
	for fast != l.last && fast.next != l.last {
		fast = fast.next.next
		slow = slow.next
	}

	if fast == l.last {
		fmt.Println(slow.val)
	}else {
		fmt.Printf("%v, %v\n", slow.val, slow.next.val)
	}

}

func (l *LinkedList[T]) HasLoop() bool {
	if l.IsEmpty(){
		return false
	}

	slow, fast := l.first, l.first
	for fast != nil && fast.next != nil {
		if fast == slow{
			return true
		}
		fast = fast.next.next
		slow = slow.next
	}
	return false
}
