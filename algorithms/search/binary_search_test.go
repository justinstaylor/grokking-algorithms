package search

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := BinarySearch(arr, 5)
	if result != 4 {
		t.Errorf("Expected 4 but got %d", result)
	}
}
