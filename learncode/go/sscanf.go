package main

import "fmt"

func main() {
	str := "12345_3248ff"

	var (
		a int64
		b string
	)
	n, err := fmt.Sscanf(str, "%d_%s", &a, &b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(a, b, n)
}

