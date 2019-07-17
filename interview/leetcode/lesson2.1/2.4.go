package main

import "fmt"

// 旋转排序数组允许重复搜索,默认升序
//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
//(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
//
//You are given a target value to search. If found in the array return true, otherwise return false.
//
//Example 1:
//
//Input: nums = [2,5,6,0,0,1,2], target = 0
//Output: true
//Example 2:
//
//Input: nums = [2,5,6,0,0,1,2], target = 3
//Output: false
func searchRotatedArr1(arr []int, key int) int {
	beg, end := 0, len(arr)-1

	for beg <= end {
		mid := (beg + end) / 2

		if arr[mid] == key {
			return mid
		}

		if arr[mid] > arr[beg] { // 左边有序
			if arr[beg] <= key && key < arr[mid] { // key 取值确实在左边
				end = mid
			} else {
				beg = mid + 1
			}
		} else if arr[mid] < arr[beg] { // 右边有序
			if arr[mid] < key && key <= arr[end] { // key 取值确实在右边
				beg = mid + 1
			} else {
				end = mid
			}
		} else {
			beg += 1

			// 处理极端情况：
			// 而如果可以有重复值，就会出现来面两种情况，[3 1 1] 和 [1 1 3 1]，对于这两种情况中间值等于最右值时，
			// 目标值3既可以在左边又可以在右边，那怎么办么，对于这种情况其实处理非常简单，只要把最右值向左一位即可继续循环，
			// 如果还相同则继续移，直到移到不同值为止
		}
	}
	return -1
}

func main() {
	//arr := []int{3, 4, 4, 5, 6, 6, 7, 7, 0, 0, 1, 2, 3}
	//re := searchRotatedArr1(arr, 4)
	//fmt.Println(arr[re])
	//re = searchRotatedArr1(arr, 3)
	//fmt.Println(arr[re])
	//re = searchRotatedArr1(arr, 7)
	//fmt.Println(arr[re])
	//re = searchRotatedArr1(arr, 6)
	//fmt.Println(arr[re])
	//re = searchRotatedArr1(arr, 0)
	//fmt.Println(arr[re])
	//re = searchRotatedArr1(arr, 9)
	//fmt.Println(re)

	arr := []int{3, 1, 1, 1, 1}
	re := searchRotatedArr1(arr, 3)
	fmt.Println(re)

	arr = []int{1, 3, 1, 1, 1}
	re = searchRotatedArr1(arr, 3)
	fmt.Println(re)

	arr = []int{1, 1, 3, 1, 1}
	re = searchRotatedArr1(arr, 3)
	fmt.Println(re)

	arr = []int{1, 1, 1, 3, 1}
	re = searchRotatedArr1(arr, 3)
	fmt.Println(re)

	arr = []int{1, 1, 1, 1, 3}
	re = searchRotatedArr1(arr, 3)
	fmt.Println(re)
}
