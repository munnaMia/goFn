package main

import (
	"fmt"

	"github.com/munnaMia/goFn/sorting"
)

func main() {
	fmt.Println("Welcome to GO FN")
	arr := []int{15, 6, 476, 75, 546}
	sorting.InsertionSort(arr)
	fmt.Println(arr)
}
