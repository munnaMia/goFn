package math

import (
	"github.com/munnaMia/goFn/utils"
)

func GCD1(x, y int) int {
	gcd := 1
	small := utils.Lowest(x, y)

	for i := gcd; i <= small; i++ {
		if x%i == 0 && y%i == 0 {
			gcd = i
		}
	}

	return gcd
}

func GCD2(x, y int) int {
	n1 := x
	n2 := y

	for n1 > 0 && n2 > 0 {
		if n1 > n2 {
			n1 = n1 % n2
		} else {
			n2 = n2 % n1
		}
	}

	if n1 == 0 {
		return n2
	}
	return n1
}
