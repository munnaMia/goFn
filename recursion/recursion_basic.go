package recursion

import "fmt"

func PrintNameNTime(n int, name string) {
	if n == 0 {
		return
	}
	fmt.Println(name)
	n--
	PrintNameNTime(n, name)
}

func PrintN(n int) {
	if n == 0 {
		return
	}
	PrintN(n - 1)
	fmt.Println(n)
}

// This is also known as forward tracing
func PrintNRev1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	PrintNRev1(n - 1)
}

// This is also known as Backward tracing
func PrintNRev3(n int) {
	if n < 1 {
		return
	}

	PrintNRev3(n - 1)

	fmt.Println(n)
}

// print i to n
func PrintNRev2(n, i int) {
	if i == n+1 {
		return
	}
	fmt.Println(i)
	i++
	PrintNRev2(n, i)
}

// Sum of n numbers
func SumOfN(n int) int {
	if n == 1 {
		return 1
	}

	return n + SumOfN(n-1)

}

// fact of n numbers
func FactoOfN(n int) int {
	if n == 1 {
		return 1
	}

	return n * SumOfN(n-1)
}

func RevArray(array []int, left, right int) {
	if left >= right {
		return
	}
	array[left], array[right] = array[right], array[left]
	left++
	right--
	RevArray(array, left, right)
}

func IsPalindromeStr(str string, i int) {
	if i >= len(str) {
		fmt.Println("Is palindrome")
		return
	}

	if str[i] != str[len(str)-1-i] {
		fmt.Println("Is not palindrome")
		return
	}

	IsPalindromeStr(str, i+1)
}

