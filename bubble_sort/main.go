package main

import "fmt"

func main() {
	test1 := []int{90, 12, 2, 34, 2, 1, 5, 3, 23, 42}
	BubbleSort(&test1)
	fmt.Println(test1)

}

func BubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
	}
}
