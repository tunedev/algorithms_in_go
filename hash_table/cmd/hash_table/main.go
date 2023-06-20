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

	fmt.Println("Most frequent",mostFrequent([]int{1,2,2,3,3,3,4}))
	fmt.Println("count unique pairs",countPairsWithDiff([]int{1,7,5,9,2,12,3}, 2))
	fmt.Println("two sums",twoSum([]int{2,7,11,15}, 9))
}

func mostFrequent(input []int) int {
	mostFrequent := 0
	freqMap := map[int]int{}

	for _, num := range input {
		value := freqMap[num]
		value += 1
		freqMap[num] = value
		if value > mostFrequent {
			mostFrequent = num
		}
	}
	return mostFrequent
}

func countPairsWithDiff(input []int, k int) int {
	set := map[int]struct{}{}
	add := func(val int) {
		if _, exist := set[val]; exist{
			return
		}
		set[val] = struct{}{}
	}

	for _, num := range input {
		add(num)
	}

	count := 0
	for num, _ := range set{
		if _, exist := set[num + k]; exist {
			count ++
		}
		if _, exist := set[num - k]; exist{
			count++
		}
		delete(set, num)
	}
	return count
}

func twoSum(input []int, target int) []int {
	seen := map[int]int{}

	for i, val := range input {
		complement := target - val
		if otherIndex, exist := seen[complement]; exist {
			return []int{otherIndex, i}
		}

		seen[val] = i
	}
	return nil
}
