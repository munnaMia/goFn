package sorting

/*
	ALGORITHM MergeSort(List)
	    n = length of List
	    IF n <= 1:
	        RETURN List

	    middle = n / 2
	    leftHalf = MergeSort(List[0...middle])
	    rightHalf = MergeSort(List[middle...n])

	    RETURN Merge(leftHalf, rightHalf)
	END ALGORITHM


	ALGORITHM Merge(Left, Right)
	    result = empty list
	    WHILE Left is not empty AND Right is not empty:
	        IF Left[0] <= Right[0]:
	            append Left[0] to result
	            remove Left[0] from Left
	        ELSE:
	            append Right[0] to result
	            remove Right[0] from Right

	    // Append remaining elements (if any)
	    append remaining Left to result
	    append remaining Right to result

	    RETURN result
	END ALGORITHM
*/

func merge(arr []int, low, mid, high int) {
	temp := make([]int, 0)
	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

func MergeSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	MergeSort(arr, low, mid)
	MergeSort(arr, mid+1, high)

	merge(arr, low, mid, high)
}

/*
	Performance Analysis
	--------------------
	Best Case Time Complexity: 		O(nlogn),"It always divides and merges, regardless of initial order."
	Average Case Time Complexity: 	O(nlogn),Consistent performance across all data distributions.
	Worst Case Time Complexity: 	O(nlogn),Guaranteed speed even on reverse-sorted data.
	Space Complexity: 				O(n),Key Trade-off: It requires extra space to hold the sub-lists.
	Stability: 						Stable,It maintains the relative order of equal elements.
*/
