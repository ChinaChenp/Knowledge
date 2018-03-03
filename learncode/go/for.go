package main

import "fmt"

func main() {
	var mp = map[string]string{
		"1": "11",
		"2": "22",
		"3": "33",
	}

	for _, v := range mp {
		fmt.Println(&v, v)
	}

	str := "chenping"
	for _, v := range str {
		fmt.Println(&v, v)
	}
}
