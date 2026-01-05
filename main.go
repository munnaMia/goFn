package main

import (
	"fmt"

	"github.com/munnaMia/goFn/sorting"
)

func main() {
	fmt.Println("Welcome to GO FN")
	arr := []int{123, 5, 3, 23, 5, 2, 234, 5, 3, 234, 234, 234, 5, 56, 5, 6, 476, 75, 546}

	fmt.Println(arr)
	sorting.BubbleSort(arr)
	fmt.Println(arr)
}
