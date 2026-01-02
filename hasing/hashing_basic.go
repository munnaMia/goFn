package hasing

import "fmt"

func CountFreq(n int) {
	hash := make(map[int]int)

	for n > 0 {
		hash[n%10]++
		n /= 10
	}

	for key, value := range hash {
		fmt.Println(key, " : ", value)
	}
}

func CountFreqHighAndLow(n int) (int, int) {
	hash := make(map[int]int)

	for n > 0 {
		hash[n%10]++
		n /= 10
	}
	high := 0
	low := 9

	for _, value := range hash {
		if high < value {
			high = value
		}
		if low > value {
			low = value
		}
	}
	return high, low
}
