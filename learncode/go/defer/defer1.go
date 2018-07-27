package main

import "fmt"

type info struct {
	b int
	c string
}

func print(a int, d *info) {
	fmt.Println(a, *d)
}

func print1(a int) {
	fmt.Println(a)
}

func  main() {
	a := 5
	d := &info{2, "chenping"}

	defer func() {
		print(a, d)
	}()

	a = 6
	d = &info{3, "chenping1"}

	for i := 0; i < 5; i++ {
		defer func() {
			print1(i)
		}()
	}
}