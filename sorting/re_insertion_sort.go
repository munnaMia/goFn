package sorting

func RecursionInsertionSort(arr []int, i int) {
	if len(arr)-1 == i {
		return
	}

	j := i
	for j > 0 && arr[j-1] > arr[j] {
		arr[j-1], arr[j] = arr[j], arr[j-1]
		j--
	}
	RecursionInsertionSort(arr, i+1)
}
