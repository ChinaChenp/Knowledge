package main

import "fmt"

func test() (int, bool) {
	return 15, true
}

func main() {
	// 是这个条件判断结构的局部变量
	num := 9
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 9 {
		fmt.Println(num, "has 1 digit")
	} else {
		if num, ok := test(); ok {
			fmt.Println("ok has multiple digits", num)
		}
		fmt.Println("has multiple digits")
	}
	fmt.Println(num, "done..............")

	a := 'A'
	fmt.Println(a)
}
