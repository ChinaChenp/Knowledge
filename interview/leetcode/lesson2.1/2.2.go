package main

import "fmt"

// 删除排序数组中重复的数字，允许重复2次
//Remove Duplicates from Sorted Array II
//􏳯􏳰
//Follow up for ”Remove Duplicates”: What if duplicates are allowed at most twice?
//For example, Given sorted array A = [1,1,1,2,2,3],
//Your function should return length = 5, and A is now [1,1,2,2,3]

const duplicates int = 2

func removeDuplicates2(arr []int) int {
	if len(arr) <= 2 {
		return len(arr)
	}

	index := duplicates
	for i := duplicates; i < len(arr); i++ {
		if arr[index-duplicates] != arr[i] {
			arr[index] = arr[i]
			index += 1
		}

		// 相等忽略
	}
	return index
}

func main() {
	arr := []int{1, 1, 1, 2, 3, 3, 3, 5, 6, 7, 7, 7}
	re := removeDuplicates2(arr)
	fmt.Println(arr[:re], re)
}
