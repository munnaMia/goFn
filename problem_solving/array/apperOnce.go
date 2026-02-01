package array

func AppearOnce(arr []int) int {
	var element int
	for _, v := range arr {
		element ^= v

	}
	return element
}
