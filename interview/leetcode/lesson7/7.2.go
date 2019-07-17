package main

import "fmt"

// 在一个排序数组中找到第一个key的index位置，没找到返回插入的index位置
// 思路：二分查找左边界和右边界
// 同7.1 思路相同，相当于std::lower_bound()

func LeftK(arr []int, key int) int {
	beg, end := 0, len(arr)-1

	mid := (end + beg) / 2
	for beg <= end {
		mid = (end + beg) / 2
		if arr[mid] > key {
			end = mid - 1
		} else if arr[mid] < key {
			beg = mid + 1
		} else { // 相等
			if mid-1 >= 0 && arr[mid-1] == arr[mid] {
				end = mid - 1
			} else {
				return mid
			}
		}
	}
	return mid
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
	fmt.Println(LeftK(arr, 14))
	fmt.Println(LeftK(arr, 4))
}
