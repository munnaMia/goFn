package array

func MoveZerosToEnd(arr []int) {
	i := 0
	j := len(arr) - 1

	for i < j {
		if arr[i] == 0 {
			if arr[j] != 0 {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				j--
			}
		} else {
			i++
		}
	}
}
