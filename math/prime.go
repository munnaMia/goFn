package math

import "math"

func IsPrime1(n int) bool {
	for i := n - 1; i >= 2; i-- {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsPrime2(n int) bool {
	if n%2 == 0 {
		return false
	}

	for i := n - 2; i >= 3; i -= 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsPrime3(n int) bool {
	k := int(math.Sqrt(float64(n)))
	for i := 2; i <= k; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// There is a more faster approch available out there learn and implement them also.