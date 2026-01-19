package main

import (
	"fmt"

	"github.com/munnaMia/goFn/problem_solving/array"
)

func main() {
	fmt.Println("Welcome to GO FN")
	arr := []int{15, 6, 476, 75, 546}
	array.RotatearraybyK(arr, 2, 'l')
	fmt.Println(arr)
}
