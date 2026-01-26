package main

import (
	"fmt"

	"github.com/munnaMia/goFn/sorting"
)

func main() {
	arr := []int{15, 6, 476, 75, 546}
	sorting.RecursiveBubbleSort(arr, len(arr))
	fmt.Println(arr)
}
