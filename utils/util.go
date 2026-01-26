package utils

func Hightest(nums ...int) int {
	high := 0
	for _, v := range nums {
		if v > high {
			high = v
		}
	}
	return high
}

func Lowest(nums ...int) int {
	low := nums[0]
	for _, v := range nums {
		if v < low {
			low = v
		}
	}
	return low
}

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}

}

func Min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

