package sorting

/*
	ALGORITHM SelectionSort(List)
    	n = length of List

    	FOR i FROM 0 TO n - 2:
        	// Assume the first unsorted element is the minimum
        	minIndex = i

        	// Scan the remaining unsorted portion to find the true minimum
        	FOR j FROM i + 1 TO n - 1:
        	    IF List[j] < List[minIndex]:
        	        minIndex = j

        	// Swap the found minimum with the first element of the unsorted part
        	IF minIndex != i:
        	    SWAP List[i] WITH List[minIndex]

    	RETURN List
	END ALGORITHM
*/

func SelectionSort(arr []int) {
	for i := range arr {
		minIdx := i

		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}

		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}

/*
	Performance Analysis
	--------------------
	Selection Sort is known for its simplicity, but it has specific performance trade-offs
	compared to Bubble Sort.

	Time Complexity
	---------------
		Best Case : O(n^2) Even if sorted, it still scans to find the minimum.
		Average Case : O(n^2) Standard nested loop behavior.
		Worst Case : O(n^2) No early exit optimization possible.

	Space Complexity : O(1) In-place algorithm; uses no extra memory.
	Auxiliary Swaps : O(n)$Key Strength: It performs a maximum of n swaps.
*/
