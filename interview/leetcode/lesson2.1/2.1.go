package main

import "fmt"

// 删除排序数组中重复的数字，只允许重复1次
// Remove Duplicates from Sorted Array
//􏳯􏳰
// Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.
// Do not allocate extra space for another array, you must do this in place with constant memory.
// For example, Given input array A = [1,1,2],
// Your function should return length = 2, and A is now [1,2].

func removeDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[index] != arr[i] {
			index += 1
			arr[index] = arr[i]
		}

	}
	return index + 1
}

func main() {
	arr := []int{1, 1, 2, 3, 3, 5, 6, 7, 7}
	re := removeDuplicates(arr)
	fmt.Println(arr[:re], re)
}
