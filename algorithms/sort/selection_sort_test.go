package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 7, 2, 4, 1, 3, 6, 9, 8, 10}
	sortedArr := SelectionSort(arr)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range sortedArr {
		if v != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], v)
		}
	}
}
