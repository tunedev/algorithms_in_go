package hashtable

import (
	"errors"
	"fmt"
	"math"

	l "github.com/tunedev/algorithms_in_go/linked_list"
)

// HashTables give us super fast lookups, below are examples of where they are useful
// spell checkers, dictionaries
// compilers used in quickly looking up address of functions and variables
// They are used to store key value pair
// Run time and common operations:
// 	insert: O(1)
// 	lookup: O(1)
// 	delete: O(1)

type Entry[T comparable] struct {
	key int
	val T
}

type HashTable[T comparable] struct{
	items []*l.LinkedList[*Entry[T]]
	size int
	loadFactor float64
	treshold int
}

var ErrKeyDoesNotExist = errors.New("key does not exist")

func NewHashTable[T comparable](arg ...int) *HashTable[T] {
	size :=  16
	if len(arg) > 0 {
		size = arg[0]
	}
	loadFactor := 0.75
	return &HashTable[T]{
		items: make([]*l.LinkedList[*Entry[T]], size),
		size: size,
		loadFactor: loadFactor,
		treshold: int(float64(size) * loadFactor),
	}
}

func (h *HashTable[T]) resize() {
	newSize := h.size * 2;
	newItems := make([]*l.LinkedList[*Entry[T]], newSize)

	for i := 0; i < h.size; i++ {
		if h.items[i] != nil {
			for _, entry := range h.items[i].ToArray() {
				newHash := h.hashWithSize(entry.key, newSize)
				if newItems[newHash] == nil {
					newItems[newHash] = l.NewLinkedList[*Entry[T]]()
				}
				newItems[newHash].AddLast(entry)
			}
		}
	}

	h.size = newSize
	h.items = newItems
	h.treshold = int(float64(newSize) * h.loadFactor)
}

func (h *HashTable[T]) Put(k int, val T) {
	if len(h.items) >= h.treshold {
		h.resize()
	}

	hash := h.hash(k)
	if h.items[hash] == nil {
		h.items[hash] = l.NewLinkedList[*Entry[T]]()
	}
	bucket := h.items[hash];
	// conirm the key is not a duplicate, to update the entry instead
	for _, entry := range bucket.ToArray() {
		if entry.key == k {
			(*entry).val = val
			return
		}
	}
	bucket.AddLast(&Entry[T]{
		key: k,
		val: val,
	})
}

func (h *HashTable[T]) Get(k int) T {
	bucket := h.items[h.hash(k)]
	if bucket != nil && !bucket.IsEmpty() {
		for _, entry := range bucket.ToArray() {
			if entry.key == k {
				return entry.val
			}
	}
	}
	return (&Entry[T]{}).val
}

func (h *HashTable[T]) Remove(k int) error {
	hash := h.hash(k)
	bucket := h.items[hash]

	if bucket == nil {
		return ErrKeyDoesNotExist
	}

	for _, entry := range bucket.ToArray() {
		if entry.key == k {
			return bucket.Remove(entry)
		}
	}
	return ErrKeyDoesNotExist
}

func (h *HashTable[T]) hash(key int) int {
	return int(math.Abs(float64(key))) % h.size
}

func (h *HashTable[T]) hashWithSize(key int, size int) int {
	return int(math.Abs(float64(key))) % size
}

func (h *HashTable[T]) String() string {
	var resp string

	for i, bucket := range h.items {
		if bucket != nil {
			resp += fmt.Sprintf("Hash: %v, contains: \n\t%v\t\n", i, bucket)
		}
	}

	return resp
}

