package main

import "fmt"

// 旋转排序数组搜索,默认升序
//Search in Rotated Sorted Array
//􏳯􏳰
//Suppose a sorted array is rotated at some pivot unknown to you beforehand.
//(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
//You are given a target value to search. If found in the array return its index, otherwise return -1. You may assume no duplicate exists in the array.

func searchRotatedArr(arr []int, key int) int {
	beg, end := 0, len(arr) - 1

	for beg <= end {
		mid := (beg + end) / 2

		if arr[mid] == key {
			return mid
		}

		if arr[mid] >= arr[beg] { // 左边有序
			if arr[beg] <= key && key <= arr[mid] { // key 取值确实在左边
				end = mid
			} else {
				beg = mid + 1
			}
		} else { //arr[mid] <= arr[beg]
			if arr[mid] <= key && key <= arr[end] { // key 取值确实在右边
				beg = mid
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func main()  {
	arr := []int{4, 5, 6, 7, 0, 1, 2, 3}
	re := searchRotatedArr(arr, 4)
	fmt.Println(arr[re])
	re = searchRotatedArr(arr, 3)
	fmt.Println(arr[re])
	re = searchRotatedArr(arr, 7)
	fmt.Println(arr[re])
	re = searchRotatedArr(arr, 6)
	fmt.Println(arr[re])
	re = searchRotatedArr(arr, 0)
	fmt.Println(arr[re])
	re = searchRotatedArr(arr, 9)
	fmt.Println(re)
}