package main

import (
	"fmt"
	"grokking-algorithms/algorithms/search"
	"grokking-algorithms/algorithms/sort"
)

func main() {
	binarySearch()
	selectionSort()
	quickSort()
	breadthFirstSearch()
}

func binarySearch() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5
	result := search.BinarySearch(arr, target)
	fmt.Printf("Found %d at index: %d\n", target, result)
}

func selectionSort() {
	arr := []int{5, 7, 2, 4, 1, 3, 6, 9, 8, 10}
	result := sort.SelectionSort(arr)
	fmt.Println("Sorted array: ", result)
}

func quickSort() {
	arr := []int{5, 7, 2, 4, 1, 3, 6, 9, 8, 10}
	result := sort.QuickSort(arr)
	fmt.Println("Sorted array: ", result)
}

func breadthFirstSearch() {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}
	start := "you"
	target := "thom"

	result := search.BreadthFirstSearch(graph, start, target)
	fmt.Printf("Found %s: %t\n", target, result)
}
