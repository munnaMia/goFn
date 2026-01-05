package sorting

func SelectionSort(arr []int) {
	for i := range arr {
		minIdx := i

		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}

		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}
