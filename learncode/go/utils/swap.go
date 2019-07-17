package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
