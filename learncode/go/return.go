package main

import (
	"fmt"
)

func SumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

func main() {
	sum, prod, diff := SumProductDiff(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)

	if x, _, _ := SumProductDiff(3, 4); x == 7 {
		fmt.Println("yes", x)
	} else {
		fmt.Println("no", x)
	}
}
