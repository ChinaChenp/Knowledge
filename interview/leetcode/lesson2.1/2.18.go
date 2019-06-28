package main

import "fmt"

// 一次只能前进一步或者二步，求n阶台阶有多少种方式

func Fibonacci1(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a + b
		//fmt.Println(a, b)
	}
	return b
}

func Fibonacci2(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return Fibonacci2(n - 1) + Fibonacci2(n - 2)
}

func main() {
	n := Fibonacci1(10)
	fmt.Println(n)

	n = Fibonacci2(10)
	fmt.Println(n)
}
