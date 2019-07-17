package main

import "fmt"

// 70. 爬楼梯
// https://leetcode-cn.com/problems/climbing-stairs/

func climbStairs(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	re := climbStairs(3)
	fmt.Println(re)
}
