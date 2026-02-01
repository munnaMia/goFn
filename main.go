package main

import (
	"fmt"

	"github.com/munnaMia/goFn/problem_solving/array"
)

func main() {
	// arr1:= []int{1,2,3,4,5,7,8}
	// arr2:= []int{1,2,3,4,7,8,9}

	arr := []int{1, 1, 1, 4,7, 1, 2, 2, 4, 4, 1, 1}

	fmt.Println(array.AppearOnce(arr))
}
