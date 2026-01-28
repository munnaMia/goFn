package array

import "fmt"

func FindMissingOne(arr []int) {
	tracker := 1
	for _, v := range arr {
		if v != tracker {
			fmt.Println("missingONe is :", tracker)
			return
		}
		tracker++
	}
}
