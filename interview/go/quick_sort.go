package main

func quickSort(arr []int, beg, end int) {
	left, right := beg, end
	base := arr[(beg+end)/2]

	for left <= right {
		for arr[right] > base && right > beg {
			right--
		}

		for arr[left] < base && left < end {
			left++
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if beg < right {
		quickSort(arr, beg, right)
	}

	if left < end {
		quickSort(arr, left, end)
	}
}
