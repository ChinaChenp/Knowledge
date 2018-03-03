package main

import "fmt"

func main() {
	str := "hello, 世界"

	for i, ch := range str {
		fmt.Println(i, ch)
	}
}
