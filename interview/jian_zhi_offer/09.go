package main

import (
	"fmt"
	"os"
	"strconv"
)

// TODO : 1）菲波那切数列 2）青蛙跳台阶问题
func Fibonacci(n int32) int64 {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci1(n int) int64 {
	base := []int64{0, 1}
	if n < 2 {
		return base[n]
	}

	first, second := base[0], base[1]
	var total int64
	for i := 0; i < n; i++ {
		total = first + second
		second = first
		first = total
	}
	return total
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	a := Fibonacci1(num)
	fmt.Println(a)
}
