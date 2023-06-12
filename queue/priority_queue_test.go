package queue

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	size := 5
	pq := NewPriorityQueue[int](size)

	// Test empty queue
	if !pq.IsEmpty() {
		t.Errorf("Queue should be empty")
	}

	// Test adding elements
	err := pq.Add(1, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = pq.Add(2, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = pq.Add(3, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	fmt.Println(pq)
	// Test removing elements
	item, err := pq.Remove()
	if err != nil || item != 1 {
		t.Errorf("Expected highest priority item 1, but got %v", item)
	}

	item, err = pq.Remove()
	if err != nil || item != 2 {
		t.Errorf("Expected second highest priority item 2, but got %v", item)
	}

	item, err = pq.Remove()
	if err != nil || item != 3 {
		t.Errorf("Expected lowest priority item 3, but got %v", item)
	}

	// Test empty queue after removals
	if !pq.IsEmpty() {
		t.Errorf("Queue should be empty")
	}

	// Test full queue
	for i := 0; i < size; i++ {
		err = pq.Add(i, i)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	if !pq.IsFull() {
		t.Errorf("Queue should be full")
	}

	// Test error when adding to full queue
	err = pq.Add(6, 1)
	if err == nil {
		t.Errorf("Expected error when adding to full queue, but got nil")
	}
}