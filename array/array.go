package array

import (
	"errors"
	"fmt"
	"math"
)

// Definitiion: An array is used to store a list of items, this items get stored sequentially in memory
// runtime complexity break down:
// 		Lookup: O(1)
// 		Insert: O(n)
// 		Delete: O(n)

var ErrIllegalArg = errors.New("illegal Argument")

type Array struct {
	items []int
	count int
}

func NewArray() *Array{
	return &Array{
		items: []int{},
	}
}

func (a *Array) String() string {
	result := ""
	for i := 0; i < a.count; i++ {
		result += fmt.Sprintf("index: %v, value: %v\n", i, a.items[i])
	}
	return result
}

func (a *Array) Insert(item int) {
	a.items = append(a.items, item)
	a.count++
}

func (a *Array) RemoveAt(index int) error {
	// Validate the index
	if index < 0 || index >= a.count {
		return ErrIllegalArg
	}
	// shift items to the left to fill the hole
	for i := index; i < a.count - 1; i++{
		a.items[i] = a.items[i + 1];
	}
	a.count--;
	return nil
}

func (a *Array) IndexOf(val int) int {
	for i := 0; i < a.count; i++{
		if a.items[i] == val {
			return i
		}
	}
	return -1
}

func (a *Array) Max() int {
	maxSoFar := math.MinInt
	for i := 0; i < a.count; i++ {
		curVal := a.items[i]
		if curVal > maxSoFar {
			maxSoFar = curVal
		}
	}
	return maxSoFar
}

func (a *Array) Intersect(comparable Array) []int {
	commonItems := NewArray()
	for i := 0; i < a.count; i++ {
		curVal := a.items[i]
		if comparable.IndexOf(curVal) > -1 && commonItems.IndexOf(curVal) == -1 {
			commonItems.Insert(curVal)
		}
	}
	return commonItems.items
}

func (a *Array) Reverse() {
	clone := make([]int, a.count)
	copy(clone, a.items)

	for i := 0; i < a.count; i++ {
		a.items[i] = clone[a.count - i - 1]
	}
}
