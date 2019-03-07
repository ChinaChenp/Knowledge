package main

import "fmt"

func add(a int) (c int, ok bool) {
	return a + 1, true
}

func test(a int) (b int, ok bool) {
	d, ok := add(a)
	fmt.Printf("%v, %v\n", d, ok)
	ok = false
	return d, ok
}

func main() {
	re, ok := test(1)
	fmt.Printf("%v, %v\n", re, ok)

	var c = 0
	a, b := 1, 2
	c, a = a, 2

	fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)

	a, b = 4, a
	fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)
}
