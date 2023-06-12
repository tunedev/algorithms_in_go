package stack

import (
	"errors"
	"fmt"
)

// Stack is a powerful data structure that can be used to solve lots of common problems
// Examples are:
// -- used in implementing undo features, navigation features(forward, backward), parsing the syntax of expressions
// Evaluate expressions.

// Runtime complexities of common stack operations
// 	Push(item) ===> O(1)
// 	Pop() ===> O(1)
// 	Peek() ===> O(1)
// 	isEmpty() ===> O(1)

type Stack[T any] struct{
	items []T
	count int
}

var ErrStackOverFlow = errors.New("Stack overflow: Stack is full")
var ErrIllegalState = errors.New("the operation could not be processed, as the state canot permit the operation")

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, size),
	}
}

func (s *Stack[T]) Push(val T) {
	if s.count == len(s.items){
		panic(ErrStackOverFlow)
	}
	s.items[s.count] = val
	s.count += 1
}

func (s *Stack[T]) Pop() T {
	if s.count == 0 {
		panic(ErrIllegalState)
	}
	s.count -= 1
	return s.items[s.count]
}

func (s *Stack[T]) Peek() T {
	if s.count == 0 {
		panic(ErrIllegalState)
	}
	return s.items[s.count - 1]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.count == 0
}

func (s *Stack[T]) String() string {
	return fmt.Sprint(s.items[0: s.count])
}