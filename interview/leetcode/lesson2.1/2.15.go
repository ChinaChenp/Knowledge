package main

import "fmt"

//Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
//For example, Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
// 思路：找到每个柱子左右两边的最大值，当前柱子水量 = min(max_left, max_right) - cur

func MaxInt(a int, b int, c ...int) int {
	out := a
	if b > a {
		out = b
	}

	for _, v := range c {
		if v > out {
			out = v
		}
	}
	return out
}

func MinInt(a int, b int, c ...int) int {
	out := a
	if b < a {
		out = b
	}

	for _, v := range c {
		if v < out {
			out = v
		}
	}
	return out
}

func trap(arr []int) int {
	if len(arr) < 1 {
		return 0
	}

	leftMax := make([]int, len(arr))
	rightMax := make([]int, len(arr))

	// 寻找每个柱子左边最大值
	for i := 1; i < len(arr); i++ {
		v := MaxInt(leftMax[i - 1], arr[i])
		leftMax[i] = v
	}

	// 寻找每个柱子右边最大值
	for j := len(arr) - 1 - 1; j > -1; j-- {
		v := MaxInt(rightMax[j + 1], arr[j])
		rightMax[j] = v
	}

	fmt.Println(leftMax)
	fmt.Println(rightMax)

	sum := 0
	for i := 0; i < len(arr); i++ {
		height := MinInt(leftMax[i], rightMax[i])
		if height > arr[i] { // 可以接水
			sum += height - arr[i]
		}
	}
	return sum
}

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(arr))
}