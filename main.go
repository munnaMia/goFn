package main

import (
	"fmt"

	"github.com/munnaMia/goFn/sorting"
)

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sorting.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
