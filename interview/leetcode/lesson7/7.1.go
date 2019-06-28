package main

import "fmt"

// 在一个可重复的排序数组中查找目标值的左边界和右边界
// 思路：二分查找左边界和右边界

func LeftK(arr []int, key int) int {
	beg, end := 0, len(arr) - 1

	for beg <= end {
		mid := (end + beg) / 2
		if arr[mid] > key {
			end = mid - 1
		} else if arr[mid] < key {
			beg = mid + 1
		} else { // 相等
			if mid - 1 >= 0 && arr[mid-1] == arr[mid] {
				end = mid - 1
			} else {
				return mid
			}
		}
	}
	return -1
}


func main() {
	arr := []int{5, 7, 7, 8, 8, 10, 13}
	fmt.Println(LeftK(arr, 7))
	fmt.Println(LeftK(arr, 8))
	fmt.Println(LeftK(arr, 5))
	fmt.Println(LeftK(arr, 10))
	fmt.Println(LeftK(arr, 11))
	fmt.Println(LeftK(arr, 13))
	fmt.Println(LeftK(arr, 9))
	fmt.Println(LeftK(arr, 13))
}