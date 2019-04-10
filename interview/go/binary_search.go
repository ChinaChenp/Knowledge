package main

import "fmt"

func binarySearch(arr []int, key int) int {
	beg, end := 0, len(arr)

	for beg <= end {
		mid := (beg + end) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			end = mid - 1
		} else {
			beg = beg + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 4, 6, 8, 9, 12, 15, 16}
	re := binarySearch(arr, 16)
	if re != -1 {
		fmt.Println(arr[re])
	} else {
		fmt.Println("not ok")
	}
}
