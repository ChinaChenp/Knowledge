package main

import "fmt"

// TODO : 输入一个整数n，求从1到n个整数的十进制表示中1出现的次数
//
func GetNumberOf1(n int) int {
	cout := 0
	for n > 0 {
		if n%10 == 1 {
			cout++
		}

		n = n / 10
	}
	return cout
}

func NumberOf1Between1AndN(n int) int {
	count := 0

	for i := 0; i <= n; i++ {
		count += GetNumberOf1(i)
	}
	return count
}

func main() {
	re := NumberOf1Between1AndN(12)
	fmt.Println(re)
}
