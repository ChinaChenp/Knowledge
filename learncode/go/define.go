package main

import "fmt"

func main() {
	one := 0
	fmt.Println(one)

	//two, one := 1, "1234567" // error
	two, one := 1, 2
	fmt.Println(one, two)
}
