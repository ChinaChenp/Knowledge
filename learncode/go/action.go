package main

import "fmt"

func main() {
	x := 1
	for x := 2; ; {
		fmt.Println(x)
		break
	}
	fmt.Println(x)

	if x := 3; true {
		fmt.Println(x)
	}
	fmt.Println(x)

}
