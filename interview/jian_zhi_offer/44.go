package main

import (
	"fmt"
	"sort"
)

// todo 判断5张扑克牌是不是连续的
func isContinuous(arr []int) bool {
	sort.Ints(arr)

	coutDiff, coutZero := 0, 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}

		coutDiff += arr[i] - arr[i-1] - 1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			coutZero++
		}
	}

	if coutDiff <= coutZero {
		return true
	}

	return false
}

func main() {
	arr := []int{3, 2, 4, 5, 7}
	re := isContinuous(arr)
	fmt.Println(re)
}
