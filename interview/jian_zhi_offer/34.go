package main

import "fmt"

// TODO :  寻找第N个丑数
func isUgly(a int) bool {
	for a%2 == 0 {
		a /= 2
	}

	for a%3 == 0 {
		a /= 3
	}

	for a%5 == 0 {
		a /= 5
	}

	if a == 1 {
		return true
	}

	return false
}

func GetnUgly(count int) int {
	var icr int
	for count > 0 {

		icr++
		if isUgly(icr) {
			count--
		}
	}

	return icr
}

func main() {
	re := GetnUgly(7)
	fmt.Println(re)
}
