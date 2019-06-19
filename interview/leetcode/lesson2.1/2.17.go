package main

import "fmt"

//给定一个非负整数组成的非空数组，在该数的基础上加一，返回一个新的数组。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//示例 1:
//
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//示例 2:
//
//输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。

func addOne(arr []int) []int {
	return addPlus(arr, 1)
}

func addPlus(arr []int, add int) []int {
	c := add
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] + c >= 10 { // 需要进位
			arr[i] = (arr[i] + c) % 10
			c = 1 // 肯定是1
		} else {
			arr[i] += c
			c = 0
		}
	}
	return arr
}

func main() {
	arr := []int{1, 4, 9, 0}
	for i := 0; i < 100; i++ {
		arr = addOne(arr)
		fmt.Println(arr)
	}
}
