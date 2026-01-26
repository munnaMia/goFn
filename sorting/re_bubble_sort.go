package sorting

func RecursiveBubbleSort(arr []int, n int) {
	if n == 1 {
		return
	}
	swaped := false
	for j := 0; j < n-1; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			swaped = true
		}
	}

	if !swaped {
		return
	}

	RecursiveBubbleSort(arr, n-1)
}
