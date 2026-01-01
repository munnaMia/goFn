package math

import (
	"math"
)

// Time Complexity : O(3/2 n) or O(n)
func Palindrome1(n int) bool {
	size := int(math.Log10(float64(n)) + 1)

	array := make([]int, size)

	idx := size - 1

	for n > 0 {
		array[idx] = n % 10
		n /= 10
		idx--
	}

	left := 0
	right := size - 1

	for left <= right {
		if array[left] != array[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func Palindrome2(n int) bool {
	rev := ReverseDigit(n)
	return rev == n
}
