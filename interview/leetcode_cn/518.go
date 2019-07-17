package main

import (
	"fmt"
	"sort"
)

// 先排序，然后左右两边向中间靠拢
func findUnsortedSubarray(arr []int) int {
	sortArr := make([]int, 0)
	sortArr = append(sortArr, arr...)

	sort.Ints(sortArr)

	beg, end := 0, len(arr)-1
	for beg < end {
		if arr[beg] == sortArr[beg] {
			beg += 1
		} else if arr[end] == sortArr[end] {
			end -= 1
		} else {
			break
		}
	}

	if beg == end { // 全部有序
		return end - beg
	}

	return end - beg + 1
}

func main() {
	arr := []int{2, 6, 4, 8, 10, 9, 15}
	re := findUnsortedSubarray(arr)
	fmt.Println(re)

	arr = []int{2, 6, 8, 10, 15}
	re = findUnsortedSubarray(arr)
	fmt.Println(re)
}
