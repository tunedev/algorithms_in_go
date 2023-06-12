package main

import (
	"fmt"

	array "github.com/tunedev/algorithms_in_go/array"
)

func main() {
	numbers := array.NewArray()
	numbers2 := array.NewArray()
	numbers.Insert(10)
	numbers2.Insert(10)
	numbers.Insert(23)
	numbers.Insert(19)
	numbers.Insert(4432)
	numbers2.Insert(4432)
	numbers.Insert(43)
	indexOf := numbers.IndexOf(4432)
	fmt.Println("indexOf =====>>>", indexOf)
	fmt.Println("maximum item =====>>>", numbers.Max())
	fmt.Println("intersect item =====>>>", numbers.Intersect(*numbers2))
	fmt.Println(numbers)
	numbers.Reverse()
	fmt.Println(numbers)
}