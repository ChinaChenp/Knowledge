package main

import (
	"fmt"
	"sort"
)

// todo : 快速寻找满足条件的两个数，无序数组
func FindNumbersWithSum(arr []int, sum int) (int, int) {
	sort.Ints(arr)
	beg, end := 0, len(arr)-1
	for beg < end {
		curSum := arr[beg] + arr[end]
		if curSum == sum {
			return arr[beg], arr[end]
		} else if curSum > sum {
			end--
		} else {
			beg++
		}
	}
	return -1, -1
}

func main() {
	arr := []int{4, 9, 3, 4, 8, 2, 1, 9}
	a, b := FindNumbersWithSum(arr, 10)
	fmt.Println(a, b)
}
