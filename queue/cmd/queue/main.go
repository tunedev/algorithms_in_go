package main

import (
	"fmt"

	queue "github.com/tunedev/algorithms_in_go/queue"
)

func main() {
	q := queue.NewPriorityQueue[string](5)
	// q.Enqueue(12)
	// q.Enqueue(121)
	// q.Enqueue(1212)
	// q.Enqueue(121212)
	// fmt.Println("INitial dequeue",q.Dequeue())
	// q.Enqueue(34)
	// fmt.Println(q)
	// fmt.Println("2nd dequeue",q.Dequeue())
	// fmt.Println("3rd dequeue",q.Dequeue())
	// fmt.Println("4th dequeue",q.Dequeue())
	// q.Enqueue(5)
	// q.Enqueue(4)
	// q.Enqueue(3)
	// fmt.Println(q)
	// fmt.Println("5th dequeue",q.Dequeue())
	// fmt.Println("6th dequeue",q.Dequeue())
	
	q.Add("3rd", 2)
	q.Add("2nd", 15)
	q.Add("2nd to the last", 1)
	q.Add("last", 0)
	q.Add("first", 119)
	fmt.Println(q)

}