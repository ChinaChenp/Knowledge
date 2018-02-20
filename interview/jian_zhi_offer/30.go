package main

import "fmt"

// TODO : 输入一个整形数组，数组里面有整数也有负数。数组中其中一个或连续多个整数组成一个
// TODO : 子序列，求最大和子序列
func FindGreatesSumSubArray(arr []int) int {
	maxSum := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		if sum <= 0 {
			sum = arr[i]
		} else {
			sum += arr[i]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	re := FindGreatesSumSubArray(arr)
	fmt.Println(re)
}
