package array

func FindUnion(arr1, arr2 []int) []int {
	hash := make(map[int]int)
	for _, v := range arr1 {
		hash[v]++
	}
	for _, v := range arr2 {
		hash[v]++
	}

	newArr := []int{}
	for key := range hash {
		newArr = append(newArr, key)
	}

	return newArr
}
