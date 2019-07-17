package main

import "fmt"

func IntMin(a, b int, num ...int) int {
	min := a
	if b < a {
		min = b
	}

	for _, v := range num {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	re := IntMin(-1, 0, 5, -5, 2, 4, 0)
	fmt.Println(re)
}
