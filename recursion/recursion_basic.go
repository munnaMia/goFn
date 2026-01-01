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
