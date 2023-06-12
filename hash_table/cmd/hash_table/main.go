package main

import (
	"fmt"

	h "github.com/tunedev/algorithms_in_go/hash_table"
)

func main() {
	table := h.NewHashTable[string](10)

	table.Put(-2, "Tunde sanusi")
	table.Put(2, "Tunedev sanusi")
	table.Put(25000, "South Prob")
	table.Put(3, "random thoughts sanusi")
	table.Put(1, "Tunedev inspirational")

	fmt.Println("first get",table.Get(-2))

	fmt.Println("second get",table.Get(2))

	fmt.Println(table)
}
