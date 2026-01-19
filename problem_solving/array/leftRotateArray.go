package array

import "fmt"

func LeftRotArrByOne1(arr []int) []int {
	temp := make([]int, 0)

	temp = append(temp, arr[1:]...)

	temp = append(temp, arr[0])

	return temp
}

func LeftRotArrByOne2(arr []int) []int {
	temp := arr[0]

	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	arr[len(arr)-1] = temp

	return arr
}

func RotatearraybyK(arr []int, k int, dir rune) {
	if k >= len(arr) {
		fmt.Println("invalid input K")
		return
	}

	temp := make([]int, 0)

	switch dir {
	case 'l':
		temp = append(temp, arr[len(arr)-k:]...)

		for i := len(arr) - k - 1; i >= 0; i-- {
			arr[k+i] = arr[i]
		}

		for i := range k {
			arr[i] = temp[i]
		}

	case 'r':
		temp = append(temp, arr[0:k]...)

		for i := 0; i < len(arr)-k; i++ {
			arr[i] = arr[i+k]
		}

		for i := range k {
			arr[len(arr)-k+i] = temp[i]
		}
	default:
		fmt.Println("Wrong direction")
	}
}
