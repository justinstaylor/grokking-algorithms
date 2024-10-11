package sort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{5, 7, 2, 4, 1, 3, 6, 9, 8, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := QuickSort(arr)

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d, but got %d", expected[i], v)
		}
	}
}
