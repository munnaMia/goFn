package sorting

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[low] // choose 1st ele as pivot

	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}

		for arr[j] > pivot && j > low {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

		fmt.Println(arr)
	}
	//swap pivot at the end low ==  pivot
	arr[low], arr[j] = arr[j], arr[low]

	return j

}

func QuickSort(arr []int, low, high int) {
	if low < high {
		// p is the partitioning index
		p := partition(arr, low, high)

		QuickSort(arr, low, p-1)
		QuickSort(arr, p+1, high)
	}
}
