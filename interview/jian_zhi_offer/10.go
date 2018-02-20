package main

import (
	"fmt"
	"os"
	"strconv"
)

func NumberOf1(n int) int {
	count := 0
	for n > 0 {
		count++
		n = (n - 1) & n
	}

	return count
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	re := NumberOf1(num)
	fmt.Println(re)
}
