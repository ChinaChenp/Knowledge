package main

import "fmt"

// TODO : 调整数组顺序使奇数位于偶数前面
func adjustOrder(arr []int, beg, end int) {

	left, right := beg, end
	for left < right {
		for right > beg && arr[right]%2 == 0 {
			right--
		}

		for left < end && arr[left]%2 == 1 {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
}

func main() {
	arr := []int{4, 0, 1, 4, 3, 7, 13, 12, 9, 5, 6, 19, 8, 3}
	length := len(arr) - 1
	adjustOrder(arr, 0, length)
	fmt.Println(arr)
}
