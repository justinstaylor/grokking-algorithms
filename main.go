package main

import (
	"fmt"
	"grokking-algorithms/algorithms/search"
	"grokking-algorithms/algorithms/sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := search.BinarySearch(arr, 5)
	fmt.Println("Found at index: ", result)

	arr = []int{5, 7, 2, 4, 1, 3, 6, 9, 8, 10}
	sortedArr := sort.SelectionSort(arr)
	fmt.Println("Sorted array: ", sortedArr)
}
