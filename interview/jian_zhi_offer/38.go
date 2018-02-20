package main

import "fmt"

func FindFirstK(arr []int, k int) int {
	length := len(arr) - 1
	beg, end := 0, length

	for beg <= end {
		mid := (beg + end) / 2
		if arr[mid] == k {
			if (mid > 0 && arr[mid-1] != k) || mid == 0 {
				return mid
			} else {
				end = mid - 1 // todo ：找第一个调整尾节点
			}
		} else if arr[mid] > k {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}
	return -1
}

func FindLastK(arr []int, k int) int {
	length := len(arr) - 1
	beg, end := 0, length

	for beg <= end {
		mid := (beg + end) / 2
		if arr[mid] == k {
			if (mid < length && arr[mid+1] != k) || mid == length {
				return mid
			} else {
				beg = mid + 1 // todo ：找最后一个调整头节点
			}
		} else if arr[mid] > k {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}
	return -1
}

func getNumersOfk(arr []int, k int) int {
	first := FindFirstK(arr, k)
	last := FindLastK(arr, k)
	fmt.Println(first, last)

	if first != -1 && last != -1 {
		return last - first + 1
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 9, 9, 9, 12, 34, 35, 57, 90}
	n := getNumersOfk(arr, 1)
	fmt.Println(n)
}
