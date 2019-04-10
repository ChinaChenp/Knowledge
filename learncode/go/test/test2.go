package main

import "fmt"

// fmt 打印指针内容

type A struct {
	a int
	b int
}

type B struct {
	c    int
	d    int
	info *A
}

func main() {
	b := &B{
		3,
		4,
		&A{
			1,
			2,
		},
	}

	//指针里面嵌套指针，第二个指针无法打印
	fmt.Printf("%v", b)
}
