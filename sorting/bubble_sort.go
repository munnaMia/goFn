package sorting

/*
	The Pseudo-Algorithm
	--------------------
	1) Input: An array (or slice) of N elements.
	2) Outer Loop: Iterate from i = 0 to N - 1.
	3) Optimization Flag: Set a boolean swapped to false.
		a) Inner Loop: Iterate from $j = 0$ to $n - i - 1$.
		b) Compare adjacent elements: if array[j] > array[j+1].
		c) Swap: if they are in the wrong order, swap them and set swapped to true.
	4) Termination: If swapped is still false after the inner loop finishes, the list is sorted—break the outer loop.
	5) Output: The sorted array.
*/

func BubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swap = true
			}
		}

		if !swap {
			return
		}

	}
}

/*
	Performance Characteristics:
	----------------------------
	Time Complexity:
		Worst/Average Case: O(n^2) — This happens when the list is in reverse order.
		Best Case: $O(n)$ — This happens if the list is already sorted

	Space Complexity:
		O(1) — It is an "in-place" algorithm, meaning it doesn't require extra memory
		proportional to the input size.Stability: It is a stable sort (it preserves the
		relative order of equal elements).
*/
