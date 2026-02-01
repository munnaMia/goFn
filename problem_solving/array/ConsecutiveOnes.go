package array

func ConsecutiveOnes(arr []int) int {
	max := 0
	counter := 0

	for idx, v := range arr {
		if v == 1 && idx != len(arr)-1 {
			counter++
		} else {
			counter++
			if counter > max {
				max = counter
				counter = 0
			}
		}
	}

	return max
}
