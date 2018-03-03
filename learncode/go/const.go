package main

import "fmt"

const (
	a = 1 << iota
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
