package math

import "math"

func CountDigit(n int) int {
	cnt := 0

	for n > 0 {
		n /= 10
		cnt++
	}

	return cnt
}

func CountDigitOptimal(n int) int {
	return int(math.Log10(float64(n))) + 1
}
