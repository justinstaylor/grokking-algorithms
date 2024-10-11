package sort

func QuickSort(arr []int) []int {
	less, more := []int{}, []int{}
	var pivot int

	if len(arr) < 2 {
		return arr
	} else {
		pivot = arr[0]
		for _, v := range arr[1:] {
			if v <= pivot {
				less = append(less, v)
			} else {
				more = append(more, v)
			}
		}
	}

	return append(append(QuickSort(less), pivot), QuickSort(more)...)
}
