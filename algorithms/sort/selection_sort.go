package sort

func SelectionSort(arr []int) []int {
	var newArr []int

	for range arr {
		smallest := FindSmallest(arr)
		newArr = append(newArr, smallest)
		arr = Remove(arr, smallest)
	}

	return newArr
}

func FindSmallest(arr []int) int {
	smallest := arr[0]

	for _, v := range arr {
		if v < smallest {
			smallest = v
		}
	}

	return smallest
}

func Remove(arr []int, item int) []int {
	index := -1
	for i, v := range arr {
		if v == item {
			index = i
			break
		}
	}

	if index == -1 {
		return arr
	}

	return append(arr[:index], arr[index+1:]...)
}
