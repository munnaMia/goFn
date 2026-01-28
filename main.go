package main

import (
	"fmt"

	"github.com/munnaMia/goFn/problem_solving/array"
)

func main() {
	arr := []int{1,2,3,4,5,7,8}
	array.FindMissingOne(arr)
	fmt.Println(arr)
}
