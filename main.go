package main

import (
	"fmt"

	searching "github.com/munnaMia/goFn/problem_solving/Searching"
)

func main() {
	fmt.Println("Welcome to GO FN")
	arr := []int{123, 5, 3, 23, 5, 2, 234, 5, 3, 234, 234, 234, 5, 56, 5, 6, 476, 75, 546}

	fmt.Println(searching.FindLargest(arr))
}
