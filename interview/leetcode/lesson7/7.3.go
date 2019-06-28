package main

import "fmt"

// 二维排序数组中查找给定值, 二维数组每行顺序递增，每行第一个数大于上一行最后一个数


func searchMatrix(arr [][]int, key int) bool {
	if len(arr) == 0 {
		return false
	}

	// 行大小
	n := len(arr)

	// 列大小
	m := len(arr[0])

	beg, end := 0, m * n - 1

	for beg <= end { // m * n 是取不到的
		mid := (beg + end) / 2
		x, y := mid / m, mid % m
		value := arr[x][y]

		if value == key {
			return true
		} else if value < key {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func main() {
	arr := [][]int {
		{1,   3,  5,  7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	fmt.Println(searchMatrix(arr, 1))
	fmt.Println(searchMatrix(arr, 50))
	fmt.Println(searchMatrix(arr, 11))
}