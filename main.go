package main

import (
	"fmt"
	"grokking-algorithms/algorithms/search"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := search.BinarySearch(arr, 5)
	fmt.Println("Found at index: ", result)
}
