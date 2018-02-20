package main

import "fmt"

// todo : 数组中只有一个数字出现过一次
func FindNumberOnce(arr []int) int {
	var re int
	for i := 0; i < len(arr); i++ {
		re ^= arr[i]
	}
	return re
}

// todo : 数组中只出现过2次的数字
func FindTwoNumberOnce(arr []int) (int, int) {
	var totalOr int
	for i := 0; i < len(arr); i++ {
		totalOr ^= arr[i]
	}

	index := leftFirstBit1(totalOr)
	fmt.Println("==> ", index, totalOr)
	first, second := 0, 0
	for i := 0; i < len(arr); i++ {
		if indexBitIs1(arr[i], uint(index)) {
			first ^= arr[i]
		} else {
			second ^= arr[i]
		}
	}
	return first, second
}

func leftFirstBit1(a int) int {
	index := 0
	for a > 0 {
		index++
		if a&1 == 1 {
			return index
		}
		a = a >> 1
	}
	return index
}

func indexBitIs1(a int, index uint) bool {
	n := a >> index
	if n&1 == 1 {
		return true
	}
	return false
}

func main() {
	arr := []int{3, 5, 7, 7, 3, 8, 5, 9, 0, 0, 9}
	re := FindNumberOnce(arr)
	fmt.Println(re)

	arr1 := []int{3, 5, 9, 9, 11, 7, 7, 3, 8, 5, 9, 0, 0, 9}
	re, re1 := FindTwoNumberOnce(arr1)
	fmt.Println(re, re1)
}
