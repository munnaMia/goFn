package array

func RemoveDuplicate(arr []int) {
	i := 0

	for j := i + 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			arr[i+1] = arr[j]
			i++
		}
	}
	for i+1 < len(arr) {
		arr[i+1] = 0
		i++
	}

}
