package main

import "fmt"

// todo : 最大公约数  f(x, y) = f(y, x%y) (y>0)
func gcd(x, y int) int {
	if y == 0 {
		return x
	}

	return gcd(y, x%y)
}


func main() {
	re := gcd(42, 30)
	fmt.Println(re)
}