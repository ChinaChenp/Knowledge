package main

import (
	"fmt"
)

func itoa(a int) string {
	var str string

	for a > 0 {
		c := '0' + a%10
		str = string(c) + str
		a /= 10
	}
	return str
}

func main() {
	re := itoa(1234567890)
	fmt.Println(re)
}