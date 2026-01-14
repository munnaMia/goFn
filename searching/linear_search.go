package searching

func LinearSearch(arr []int, ele int) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}

	return false
}
