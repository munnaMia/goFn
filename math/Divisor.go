package math

import "math"

func Divisor1(n int) []int {
	arr := make([]int, 0)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}

	return arr
}

func Divisor2(n int) []int {
	arr := make([]int, 0)

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			arr = append(arr, i)

			if i != n/i {
				arr = append(arr, n/i)
			}
		}
	}

	return arr
}
